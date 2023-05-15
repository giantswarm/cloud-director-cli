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
	"net/url"

	"github.com/vmware/cloud-provider-for-cloud-director/pkg/vcdsdk"
	"github.com/vmware/go-vcloud-director/v2/types/v56"
)

type VappManager struct {
	Client *vcdsdk.Client
}

func (manager *VappManager) List() []*types.QueryResultVAppRecordType {
	filter := "vdc==" + url.QueryEscape(manager.Client.VDC.Vdc.HREF)
	notEncodedParams := map[string]string{"type": "vApp", "filter": filter, "filterEncoded": "true"}
	results, err := manager.Client.VDC.QueryWithNotEncodedParams(nil, notEncodedParams)
	if err != nil {
		log.Fatal(err)
	}
	return results.Results.VAppRecord
}

func (manager *VappManager) Delete(names []string) {
	m, err := vcdsdk.NewVDCManager(manager.Client, "", "")
	if err != nil {
		log.Fatal(err)
	}

	for _, name := range names {
		err = m.DeleteVApp(name)
		if err != nil {
			log.Fatal(err)
		}
	}
}
