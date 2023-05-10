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

	"github.com/vmware/cloud-provider-for-cloud-director/pkg/vcdsdk"
	"github.com/vmware/go-vcloud-director/v2/types/v56"
)

type VirtualServiceManager struct {
	Client *vcdsdk.Client
}

type VirtualServiceListParams struct {
	Network string
}

func (manager *VirtualServiceManager) List(params VirtualServiceListParams) []*types.NsxtAlbVirtualService {
	gateway := getGatewayManager(manager.Client, params.Network)
	vSvcs, err := manager.Client.VCDClient.GetAllAlbVirtualServices(gateway.GatewayRef.Id, nil)
	if err != nil {
		log.Fatal(err)
	}

	result := make([]*types.NsxtAlbVirtualService, len(vSvcs))
	for i, vs := range vSvcs {
		result[i] = vs.NsxtAlbVirtualService
	}

	return result
}

func getGatewayManager(c *vcdsdk.Client, network string) *vcdsdk.GatewayManager {
	if network == "" {
		nw, err := c.VDC.GetNetworkList()
		if err != nil {
			log.Fatal(err)
		}
		if nw == nil || len(nw) == 0 {
			log.Fatal(fmt.Errorf("no networks detected"))
		}
		network = nw[0].Name
	}

	gateway, err := vcdsdk.NewGatewayManager(context.Background(), c, network, "")
	if err != nil {
		log.Fatal(err)
	}
	return gateway
}

func (manager *VirtualServiceManager) Delete(names []string, failIfAbsent bool, network string) {
	gateway := getGatewayManager(manager.Client, network)
	for _, vs := range names {
		err := gateway.DeleteVirtualService(context.Background(), vs, failIfAbsent)
		if err != nil {
			log.Fatal(err)
		}
	}
}
