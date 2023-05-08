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
	"github.com/giantswarm/cloud-director-cli/pkg/vcd/client"
	"github.com/giantswarm/cloud-director-cli/pkg/vcd/utils"
	"log"

	"github.com/vmware/go-vcloud-director/v2/govcd"
	"github.com/vmware/go-vcloud-director/v2/types/v56"
)

func ListAports(items bool) []*govcd.NsxtAppPortProfile {
	cache := client.Cache{}
	c, e := cache.CachedClient(items)
	if e != nil {
		log.Fatal(e)
	}
	//gateway := getGatewayManager(c)
	//gateway.DeleteAppPortProfile()
	org, err := c.VCDClient.GetOrgByName(c.ClusterOrgName)
	if err != nil {
		log.Fatal(err)
	}
	aports, err := org.GetAllNsxtAppPortProfiles(nil, types.ApplicationPortProfileScopeTenant)
	if err != nil {
		log.Fatal(err)
	}
	return aports
}

func DeleteAport(names []string, failIfAbsent bool, yes bool, verbose bool, network string) {
	if len(names) == 0 {
		log.Fatal("Provide at least 1 name of a Application Port Profile")
	}
	cache := client.Cache{}
	c, e := cache.CachedClient(verbose)
	if e != nil {
		log.Fatal(e)
	}
	gateway := getGatewayManager(c, network)
	if !yes {
		fmt.Printf("Are you sure you want to delete following Application Port Profiles: %v [y/n]?\n", names)
		var char rune
		_, err := fmt.Scanf("%c", &char)
		if err != nil {
			log.Fatal(err)
		}
		if char != 'y' && char != 'Y' {
			return
		}
	}
	for _, a := range names {
		err := gateway.DeleteAppPortProfile(a, failIfAbsent)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func PrintAports(output string, verbose bool) {
	items := ListAports(verbose)
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
