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
	"github.com/vmware/cloud-provider-for-cloud-director/pkg/vcdsdk"
	"github.com/vmware/go-vcloud-director/v2/types/v56"
	"log"
)

type LoadBalancerPoolManager struct {
	Client *vcdsdk.Client
}

type LBListParams struct {
	Network string
}

func (manager *LoadBalancerPoolManager) List(params LBListParams) []*types.NsxtAlbPool {
	gateway := getGatewayManager(manager.Client, params.Network)
	lbPools, err := manager.Client.VCDClient.GetAllAlbPools(gateway.GatewayRef.Id, nil)
	if err != nil {
		log.Fatal(err)
	}

	result := make([]*types.NsxtAlbPool, len(lbPools))
	for i, lbPool := range lbPools {
		result[i] = lbPool.NsxtAlbPool
	}

	return result
}

func (manager *LoadBalancerPoolManager) Delete(names []string, failIfAbsent bool, verbose bool, cascade bool, network string) {
	gateway := getGatewayManager(manager.Client, network)

	for _, lb := range names {
		if cascade {
			albPool, err := manager.Client.VCDClient.GetAlbPoolByName(gateway.GatewayRef.Id, lb)
			if err != nil {
				log.Fatal(fmt.Errorf("unable to get load balancer pool [%s]: [%v]", lb, err))
			}
			for _, vssRef := range albPool.NsxtAlbPool.VirtualServiceRefs {
				fmt.Printf("Cascading virtual service:[%s]\n", vssRef.Name)
				err = gateway.DeleteVirtualService(context.Background(), vssRef.Name, false)
				if err != nil {
					log.Fatal(fmt.Errorf("unable to delete virtual service [%s]: [%v]", vssRef.Name, err))
				}
			}
		}

		fmt.Printf("Deleting load balancer pool:[%s]\n", lb)
		err := gateway.DeleteLoadBalancerPool(context.Background(), lb, failIfAbsent)
		if err != nil {
			log.Fatal(fmt.Errorf("unable to delete load balancer pool [%s]: [%v]", lb, err))
		}
	}
}
