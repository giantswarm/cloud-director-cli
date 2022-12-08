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
	"github.com/vmware/go-vcloud-director/v2/govcd"
	"log"
)

func ListLBPools(verboseClient bool) []*govcd.NsxtAlbPool {
	cache := Cache{}
	c, e := cache.CachedClient(verboseClient)
	if e != nil {
		log.Fatal(e)
	}
	gateway := getGatewayManager(c)
	lbPoos, err := c.VCDClient.GetAllAlbPools(gateway.GatewayRef.Id, nil)
	if err != nil {
		log.Fatal(err)
	}

	return lbPoos
}

func DeleteLBPool(names []string, failIfAbsent bool, yes bool, verboseClient bool) error {
	if len(names) == 0 {
		log.Fatal("Provide at least 1 name of a LB Pool")
	}
	cache := Cache{}
	c, e := cache.CachedClient(verboseClient)
	if e != nil {
		log.Fatal(e)
	}
	gateway := getGatewayManager(c)
	if !yes {
		fmt.Printf("Are you sure you want to delete following LB Pools: %v [y/n]?\n", names)
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
		err := gateway.DeleteLoadBalancerPool(context.Background(), vs, failIfAbsent)
		if err != nil {
			log.Fatal(err)
		}
	}
	return nil
}

func PrintLBPools(verbose bool, verboseClient bool) error {
	var headerPrinted bool
	for _, lbpool := range ListLBPools(verboseClient) {
		if !verbose {
			fmt.Println(lbpool.NsxtAlbPool.Name)
		} else {
			if !headerPrinted {
				fmt.Printf("%-90s\t%-17s\t%-14s\t\n", "NAME", "ALGOTITHM", "MEMBERS")
				headerPrinted = true
			}
			l := lbpool.NsxtAlbPool
			fmt.Printf("%-90s\t%-17s\t%v\t\n", l.Name, l.Algorithm, l.MemberCount)
		}
	}
	return nil
}
