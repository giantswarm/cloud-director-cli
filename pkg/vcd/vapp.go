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

	"github.com/vmware/cloud-provider-for-cloud-director/pkg/vcdsdk"
	"github.com/vmware/go-vcloud-director/v2/types/v56"

	"github.com/giantswarm/cloud-director-cli/pkg/vcd/utils"
)

type VappManager struct {
	Client *vcdsdk.Client
}

func (manager *VappManager) List() []*types.ResourceReference {
	vapps := manager.Client.VDC.GetVappList()
	return vapps
}

func (manager *VappManager) Delete(names []string) {
	m, err := vcdsdk.NewVDCManager(manager.Client, "", "")
	if err != nil {
		log.Fatal(err)
	}

	for _, name := range names {
		err2 := m.DeleteVApp(name)
		if err2 != nil {
			log.Fatal(err2)
		}
	}
}

func (manager *VappManager) Print(output string, items []*types.ResourceReference) {
	switch output {
	case "json":
		utils.PrintJson(items)
	case "yaml":
		utils.PrintYaml(items)
	default:
		var headerPrinted bool
		for _, vapp := range items {
			if output == "names" {
				fmt.Println(vapp.Name)
			} else {
				if !headerPrinted {
					fmt.Printf("%-35s\t%-16s\t\n", "NAME", "ID")
					headerPrinted = true
				}
				fmt.Printf("%-35s\t%-16s\t\n", vapp.Name, vapp.ID)
			}
		}
	}
}
