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
	"github.com/vmware/go-vcloud-director/v2/types/v56"
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
	vsManager := VirtualServiceManager{
		Client: manager.Client,
	}

	for _, lb := range names {
		fmt.Printf("Deleting load balancer pool:[%s]\n", lb)
		err := gateway.DeleteLoadBalancerPool(context.Background(), lb, failIfAbsent)
		if err != nil {
			if strings.Contains(err.Error(), "obtained [400]") {
				if !cascade {
					fmt.Fprintf(os.Stderr, "First delete the associated virtual service\n")
					log.Fatal(err)
				}
				// try to delete the associated virtual services first and then re-try
				fmt.Printf("Trying to delete the Virtual Services %v first\n", names)
				vsManager.Delete(names, failIfAbsent, network)
				manager.Delete(names, failIfAbsent, verbose, false, network)
			} else {
				log.Fatal(err)
			}
		}
	}
}
