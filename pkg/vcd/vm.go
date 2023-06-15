/*
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package vcd

import (
	"fmt"
	"log"

	"github.com/vmware/go-vcloud-director/v2/govcd"

	"github.com/giantswarm/cloud-director-cli/pkg/vcd/utils"

	"github.com/vmware/cloud-provider-for-cloud-director/pkg/vcdsdk"
	"github.com/vmware/go-vcloud-director/v2/types/v56"
)

type VmManager struct {
	Client *vcdsdk.Client
}

type VmListParams struct {
	Vapp         string
	OnlyTemplate bool
}

func (manager *VmManager) List(params VmListParams) []*types.QueryResultVMRecordType {
	var filter types.VmQueryFilter
	if params.OnlyTemplate {
		filter = types.VmQueryFilterOnlyTemplates
	} else {
		filter = types.VmQueryFilterOnlyDeployed
	}
	vms, err := manager.Client.VDC.QueryVmList(filter)
	if err != nil {
		log.Fatal(err)
	}

	if params.Vapp == "" {
		return vms
	} else {
		return utils.Filter(vms, func(recordType *types.QueryResultVMRecordType) bool {
			return recordType.ContainerName == params.Vapp
		})
	}
}

func (manager *VmManager) Delete(names []string, vapp string) {
	vApp, err := manager.Client.VDC.GetVAppByName(vapp, true)
	if err != nil {
		log.Fatal(fmt.Errorf("unable to find vApp from name [%s]: [%v]", vapp, err))
	}

	for _, vmName := range names {
		vm, err := vApp.GetVMByName(vmName, true)
		if err != nil {
			log.Fatal(fmt.Errorf("unable to get vm [%s] in vApp [%s]: [%v]", vmName, vmName, err))
		}

		ensureThereIsNoAttachedDisk(vm)

		if err = vm.Delete(); err != nil {
			log.Fatal(fmt.Errorf("unable to delete vm [%s] in vApp [%s]: [%v]", vmName, vapp, err))
		}
	}
}

// inspired from https://github.com/giantswarm/cluster-api-provider-cloud-director/blob/a40b68e4b395ed04edb24c8e3b6e0e11cd9d4087/controllers/vcdmachine_controller.go#L1261
func ensureThereIsNoAttachedDisk(vm *govcd.VM) {
	if vm.VM.VmSpecSection != nil && vm.VM.VmSpecSection.DiskSection != nil {
		for _, diskSettings := range vm.VM.VmSpecSection.DiskSection.DiskSettings {
			if diskSettings.Disk != nil {
				log.Fatal(fmt.Sprintf("Cannot delete VM when there is an attached disk. vm:[%s] disk:[%s]",
					vm.VM.Name, diskSettings.Disk.Name))
			}
		}
	}
}
