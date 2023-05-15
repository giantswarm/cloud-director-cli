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
	"log"

	"github.com/vmware/cloud-provider-for-cloud-director/pkg/vcdsdk"

	"github.com/vmware/go-vcloud-director/v2/types/v56"
)

type AppPortManager struct {
	Client *vcdsdk.Client
}

func (manager *AppPortManager) List() []*types.NsxtAppPortProfile {
	org, err := manager.Client.VCDClient.GetOrgByName(manager.Client.ClusterOrgName)
	if err != nil {
		log.Fatal(err)
	}
	aports, err := org.GetAllNsxtAppPortProfiles(nil, types.ApplicationPortProfileScopeTenant)
	if err != nil {
		log.Fatal(err)
	}

	result := make([]*types.NsxtAppPortProfile, len(aports))
	for i, aport := range aports {
		result[i] = aport.NsxtAppPortProfile
	}

	return result
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
