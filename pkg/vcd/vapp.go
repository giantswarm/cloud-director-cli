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
	"log"
	"net/url"
	"strconv"

	"github.com/vmware/cloud-provider-for-cloud-director/pkg/vcdsdk"
	"github.com/vmware/go-vcloud-director/v2/types/v56"

	"github.com/giantswarm/cloud-director-cli/pkg/vcd/utils"
)

type VappManager struct {
	Client *vcdsdk.Client
}

func (manager *VappManager) List() []*types.QueryResultVAppRecordType {
	filter := "vdc==" + url.QueryEscape(manager.Client.VDC.Vdc.HREF)

	results := make([]*types.QueryResultVAppRecordType, 0)
	page := 1
	for {
		notEncodedParams := map[string]string{"type": "vApp", "filter": filter, "filterEncoded": "true", "page": strconv.Itoa(page)}
		pageResult, err := manager.Client.VDC.QueryWithNotEncodedParams(nil, notEncodedParams)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, pageResult.Results.VAppRecord...)
		page++

		if !utils.HasNextPageLink(pageResult.Results.Link) {
			break
		}
	}
	return results
}

func (manager *VappManager) Delete(names []string) {
	m, err := vcdsdk.NewVDCManager(manager.Client, "", "")
	if err != nil {
		log.Fatal(err)
	}

	for _, name := range names {
		fmt.Printf("Deleting vApp:[%s]\n", name)
		err = m.DeleteVApp(name)
		if err != nil {
			log.Fatal(err)
		}
	}
}
