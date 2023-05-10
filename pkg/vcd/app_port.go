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

	"github.com/giantswarm/cloud-director-cli/pkg/vcd/utils"

	"github.com/vmware/go-vcloud-director/v2/govcd"
	"github.com/vmware/go-vcloud-director/v2/types/v56"
)

type AppPortManager struct {
	Client *vcdsdk.Client
}

func (manager *AppPortManager) List() []*govcd.NsxtAppPortProfile {
	org, err := manager.Client.VCDClient.GetOrgByName(manager.Client.ClusterOrgName)
	if err != nil {
		log.Fatal(err)
	}
	aports, err := org.GetAllNsxtAppPortProfiles(nil, types.ApplicationPortProfileScopeTenant)
	if err != nil {
		log.Fatal(err)
	}
	return aports
}

func (manager *AppPortManager) Delete(names []string, failIfAbsent bool, network string) {
	gateway := getGatewayManager(manager.Client, network)
	for _, a := range names {
		err := gateway.DeleteAppPortProfile(a, failIfAbsent)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (manager *AppPortManager) Print(output string, items []*govcd.NsxtAppPortProfile) {
	switch output {
	case "json":
		utils.PrintJson(items)
	case "yaml":
		utils.PrintYaml(items)
	default:
		var headerPrinted bool
		for _, aport := range items {
			if output == "names" {
				fmt.Println(aport.NsxtAppPortProfile.Name)
			} else {
				if !headerPrinted {
					fmt.Printf("%-110s\t%-8s\t%-14s\t\n", "NAME", "PROTOCOL", "PORTS")
					headerPrinted = true
				}
				a := aport.NsxtAppPortProfile
				protocol := "unknown"
				port := "unknown"
				if len(a.ApplicationPorts) > 0 {
					protocol = a.ApplicationPorts[0].Protocol
					if len(a.ApplicationPorts[0].DestinationPorts) > 0 {
						port = a.ApplicationPorts[0].DestinationPorts[0]
					}
				}
				fmt.Printf("%-110s\t%-8s\t%s\t\n", a.Name, protocol, port)
			}
		}
	}
}
