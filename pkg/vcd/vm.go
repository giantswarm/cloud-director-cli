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
	"log"

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
	m, err := vcdsdk.NewVDCManager(manager.Client, "", "")
	if err != nil {
		log.Fatal(err)
	}
	for _, vm := range names {
		err = m.DeleteVM(vapp, vm)
		if err != nil {
			log.Fatal(err)
		}
	}
}
