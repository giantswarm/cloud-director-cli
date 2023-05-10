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
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/vmware/cloud-provider-for-cloud-director/pkg/vcdsdk"

	"github.com/giantswarm/cloud-director-cli/pkg/vcd/utils"

	"github.com/vmware/go-vcloud-director/v2/govcd"
)

type LoadBalancerPoolManager struct {
	Client *vcdsdk.Client
}

func (manager *LoadBalancerPoolManager) List(network string) []*govcd.NsxtAlbPool {
	gateway := getGatewayManager(manager.Client, network)
	lbPoos, err := manager.Client.VCDClient.GetAllAlbPools(gateway.GatewayRef.Id, nil)
	if err != nil {
		log.Fatal(err)
	}

	return lbPoos
}

func (manager *LoadBalancerPoolManager) Delete(names []string, failIfAbsent bool, verbose bool, cascade bool, network string) {
	gateway := getGatewayManager(manager.Client, network)
	vsManager := VirtualServiceManager{
		Client: manager.Client,
	}

	for _, lb := range names {
		err := gateway.DeleteLoadBalancerPool(context.Background(), lb, failIfAbsent)
		if err != nil {
			if strings.Contains(err.Error(), "obtained [400]") {
				if !cascade {
					fmt.Fprintf(os.Stderr, "First delete the associated virtual service\n")
					log.Fatal(err)
				}
				// try to delete the associated virtual services first and then re-try
				fmt.Printf("Trying to delete the Virtual Services %v first\n", names)
				vsManager.Delete(names, failIfAbsent, verbose, network)
				manager.Delete(names, failIfAbsent, verbose, false, network)
			} else {
				log.Fatal(err)
			}
		}
	}
}

func (manager *LoadBalancerPoolManager) Print(output string, items []*govcd.NsxtAlbPool) {
	switch output {
	case "json":
		utils.PrintJson(items)
	case "yaml":
		utils.PrintYaml(items)
	default:
		var headerPrinted bool
		for _, lbpool := range items {
			if output == "names" {
				fmt.Println(lbpool.NsxtAlbPool.Name)
			} else {
				if !headerPrinted {
					fmt.Printf("%-90s\t%-17s\t%-9s\t%s\t\n", "NAME", "ALGOTITHM", "MEMBERS", "ENABLED")
					headerPrinted = true
				}
				l := lbpool.NsxtAlbPool
				fmt.Printf("%-90s\t%-17s\t%-9v\t%t\t\n", l.Name, l.Algorithm, l.MemberCount, *l.Enabled)
			}
		}
	}
}
