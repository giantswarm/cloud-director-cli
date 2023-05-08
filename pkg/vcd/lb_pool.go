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
	"github.com/giantswarm/cloud-director-cli/pkg/vcd/client"
	"github.com/giantswarm/cloud-director-cli/pkg/vcd/utils"
	"log"
	"os"
	"strings"

	"github.com/vmware/go-vcloud-director/v2/govcd"
)

func ListLBPools(verboseClient bool, network string) []*govcd.NsxtAlbPool {
	cache := client.Cache{}
	c, e := cache.CachedClient(verboseClient)
	if e != nil {
		log.Fatal(e)
	}
	gateway := getGatewayManager(c, network)
	lbPoos, err := c.VCDClient.GetAllAlbPools(gateway.GatewayRef.Id, nil)
	if err != nil {
		log.Fatal(err)
	}

	return lbPoos
}

func DeleteLBPool(names []string, failIfAbsent bool, verbose bool, cascade bool, network string) {
	if len(names) == 0 {
		log.Fatal("Provide at least 1 name of a LB Pool")
	}
	cache := client.Cache{}
	c, e := cache.CachedClient(verbose)
	if e != nil {
		log.Fatal(e)
	}
	gateway := getGatewayManager(c, network)
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
				DeleteVs(names, failIfAbsent, verbose, network)
				DeleteLBPool(names, failIfAbsent, verbose, false, network)
			} else {
				log.Fatal(err)
			}
		}
	}
}

func PrintLBPools(output string, verbose bool, network string) {
	items := ListLBPools(verbose, network)
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
