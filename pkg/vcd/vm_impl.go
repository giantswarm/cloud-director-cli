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
	"github.com/vmware/go-vcloud-director/v2/types/v56"
	"log"
)

func ListVMs(verbose bool, onlyTemplates bool) []*types.QueryResultVMRecordType {
	cache := Cache{}
	c, e := cache.CachedClient(verbose)
	if e != nil {
		log.Fatal(e)
	}
	var filter types.VmQueryFilter
	if onlyTemplates {
		filter = types.VmQueryFilterOnlyTemplates
	} else {
		filter = types.VmQueryFilterOnlyDeployed
	}
	vms, err := c.VDC.QueryVmList(filter)
	if err != nil {
		log.Fatal(err)
	}
	return vms
}

func PrintVMs(verbose bool, onlyTemplates bool, vapp string) error {
	var headerPrinted bool
	for _, vm := range ListVMs(verbose, onlyTemplates) {
		if vapp != "" && vm.ContainerName != vapp {
			continue
		}
		if !verbose {
			fmt.Println(vm.Name)
		} else {
			if !headerPrinted {
				fmt.Printf("\n\n%-35s\t%-16s\t%-10s\t%s\t", "NAME", "VAPP", "STATUS", "DEPLOYED")
				headerPrinted = true
			}
			fmt.Printf("\n%-35s\t%-16s\t%-10s\t%t\t", vm.Name, vm.ContainerName, vm.Status, vm.Deployed)
		}
	}
	return nil
}
