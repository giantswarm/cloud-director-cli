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

	"github.com/vmware/cloud-provider-for-cloud-director/pkg/vcdsdk"
	"github.com/vmware/go-vcloud-director/v2/types/v56"
)

type DiskManager struct {
	Client *vcdsdk.Client
}

type DiskListParams struct {
	Unattached bool
}

func (manager *DiskManager) List(params DiskListParams) []*types.DiskRecordType {
	filter := "vdc==" + url.QueryEscape(manager.Client.VDC.Vdc.HREF)
	if params.Unattached {
		filter = filter + ",attachedVmCount==0"
	}
	notEncodedParams := map[string]string{"type": "disk", "filter": filter, "filterEncoded": "true"}
	results, err := manager.Client.VDC.QueryWithNotEncodedParams(nil, notEncodedParams)
	if err != nil {
		log.Fatal(err)
	}

	return results.Results.DiskRecord
}

func (manager *DiskManager) Delete(names []string) {
	for _, name := range names {
		disks, err := manager.Client.VDC.GetDisksByName(name, false)
		if err != nil {
			log.Fatal(err)
		}
		for _, d := range *disks {
			t, e := d.Delete()
			if e != nil {
				log.Fatal(e)
			}
			fmt.Printf("delete task was created: %s\n", t.Task.ID)
		}
	}
}
