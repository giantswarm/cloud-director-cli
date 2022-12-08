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
	"context"
	"encoding/json"
	"fmt"
	"github.com/vmware/cloud-provider-for-cloud-director/pkg/vcdsdk"
	"github.com/vmware/go-vcloud-director/v2/govcd"
	"log"
)

func ListVs(verboseClient bool) []*govcd.NsxtAlbVirtualService {
	cache := Cache{}
	c, e := cache.CachedClient(verboseClient)
	if e != nil {
		log.Fatal(e)
	}
	gateway := getGatewayManager(c)
	vSvcs, err := c.VCDClient.GetAllAlbVirtualServices(gateway.GatewayRef.Id, nil)
	if err != nil {
		log.Fatal(err)
	}

	return vSvcs
}

func getGatewayManager(c *vcdsdk.Client) *vcdsdk.GatewayManager {
	nw, err := c.VDC.GetNetworkList()
	if err != nil {
		log.Fatal(err)
	}
	if nw == nil || len(nw) == 0 {
		log.Fatal(fmt.Errorf("no networks detected"))
	}
	gateway, err := vcdsdk.NewGatewayManager(context.Background(), c, nw[0].Name, "")
	if err != nil {
		log.Fatal(err)
	}
	return gateway
}

func DeleteVs(names []string, failIfAbsent bool, yes bool, verboseClient bool) error {
	if len(names) == 0 {
		log.Fatal("Provide at least 1 name of a Virtual Service")
	}
	cache := Cache{}
	c, e := cache.CachedClient(verboseClient)
	if e != nil {
		log.Fatal(e)
	}
	gateway := getGatewayManager(c)
	if !yes {
		fmt.Printf("Are you sure you want to delete following Virtual Services: %v [y/n]?\n", names)
		var char rune
		_, err := fmt.Scanf("%c", &char)
		if err != nil {
			log.Fatal(err)
		}
		if char != 'y' && char != 'Y' {
			return nil
		}
	}
	for _, vs := range names {
		err := gateway.DeleteVirtualService(context.Background(), vs, failIfAbsent)
		if err != nil {
			log.Fatal(err)
		}
	}
	return nil
}

func PrintVs(output string, verboseClient bool) error {
	var headerPrinted bool
	if output == "json" {
		j, err := json.Marshal(ListVs(verboseClient))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(j))
		return nil
	}
	for _, svc := range ListVs(verboseClient) {
		if output == "names" {
			fmt.Println(svc.NsxtAlbVirtualService.Name)
		} else {
			if !headerPrinted {
				fmt.Printf("%-90s\t%-17s\t%-14s\t\n", "NAME", "IP", "HEALTH")
				headerPrinted = true
			}
			s := svc.NsxtAlbVirtualService
			fmt.Printf("%-90s\t%-17s\t%v\t\n", s.Name, s.VirtualIpAddress, s.HealthStatus)
		}
	}
	return nil
}
