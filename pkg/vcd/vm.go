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

	"github.com/giantswarm/cloud-director-cli/pkg/vcd/utils"

	"github.com/vmware/cloud-provider-for-cloud-director/pkg/vcdsdk"
	"github.com/vmware/go-vcloud-director/v2/types/v56"
)

type VmManager struct {
	Client *vcdsdk.Client
}

func (manager *VmManager) List(onlyTemplates bool) []*types.QueryResultVMRecordType {
	var filter types.VmQueryFilter
	if onlyTemplates {
		filter = types.VmQueryFilterOnlyTemplates
	} else {
		filter = types.VmQueryFilterOnlyDeployed
	}
	vms, err := manager.Client.VDC.QueryVmList(filter)
	if err != nil {
		log.Fatal(err)
	}
	return vms
}

func (manager *VmManager) Delete(names []string, vapp string, verbose bool) {
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

func (manager *VmManager) Print(outputFormat string, items []*types.QueryResultVMRecordType, vapp string) {
	switch outputFormat {
	case "json":
		utils.PrintJson(items)
	case "yaml":
		utils.PrintYaml(items)
	default:
		var headerPrinted bool
		for _, vm := range items {
			if vapp != "" && vm.ContainerName != vapp {
				continue
			}
			if outputFormat == "names" {
				fmt.Println(vm.Name)
			} else {
				if !headerPrinted {
					fmt.Printf("%-35s\t%-16s\t%-10s\t%-8s\t%-16s\t\n", "NAME", "VAPP", "STATUS", "DEPLOYED", "IP")
					headerPrinted = true
				}
				fmt.Printf("%-35s\t%-16s\t%-10s\t%-8t\t%-16s\t\n", vm.Name, vm.ContainerName, vm.Status, vm.Deployed, vm.IpAddress)
			}
		}
	}
}
