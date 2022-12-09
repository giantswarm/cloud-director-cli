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
	"github.com/vmware/cloud-provider-for-cloud-director/pkg/vcdsdk"
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

func DeleteVMs(names []string, vapp string, yes bool, verbose bool) {
	if len(names) == 0 {
		log.Fatal("Provide at least 1 name of a VM")
	}
	cache := Cache{}
	c, e := cache.CachedClient(verbose)
	if e != nil {
		log.Fatal(e)
	}
	m, err := vcdsdk.NewVDCManager(c, "", "")
	if err != nil {
		log.Fatal(err)
	}
	if !yes {
		fmt.Printf("Are you sure you want to delete following VMs: %v [y/n]?\n", names)
		var char rune
		_, err := fmt.Scanf("%c", &char)
		if err != nil {
			log.Fatal(err)
		}
		if char != 'y' && char != 'Y' {
			return
		}
	}
	for _, vm := range names {
		m.DeleteVM(vapp, vm)
	}
}

func PrintVMs(output string, verbose bool, onlyTemplates bool, vapp string) {
	items := ListVMs(verbose, onlyTemplates)
	switch output {
	case "json":
		PrintJson(items)
	case "yaml":
		PrintYaml(items)
	default:
		var headerPrinted bool
		for _, vm := range items {
			if vapp != "" && vm.ContainerName != vapp {
				continue
			}
			if output == "names" {
				fmt.Println(vm.Name)
			} else {
				if !headerPrinted {
					fmt.Printf("%-35s\t%-16s\t%-10s\t%s\t\n", "NAME", "VAPP", "STATUS", "DEPLOYED")
					headerPrinted = true
				}
				fmt.Printf("%-35s\t%-16s\t%-10s\t%t\t\n", vm.Name, vm.ContainerName, vm.Status, vm.Deployed)
			}
		}
	}
}
