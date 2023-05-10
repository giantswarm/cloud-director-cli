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

	"github.com/giantswarm/cloud-director-cli/pkg/vcd/utils"
)

type DiskManager struct {
	Client *vcdsdk.Client
}

func (manager *DiskManager) List() []*types.DiskRecordType {
	filter := "vdc==" + url.QueryEscape(manager.Client.VDC.Vdc.HREF)
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

func (manager *DiskManager) Print(outputFormat string, unattached bool, items []*types.DiskRecordType) {
	switch outputFormat {
	case "json":
		utils.PrintJson(items)
	case "yaml":
		utils.PrintYaml(items)
	default:
		var headerPrinted bool
		for _, d := range items {
			if unattached && d.AttachedVmCount > 0 {
				continue
			}
			if outputFormat == "names" {
				fmt.Println(d.Name)
			} else {
				if !headerPrinted {
					fmt.Printf("%-45s\t%-10s\t%-10s\t%s\t%-10s\t\n", "NAME", "SIZE(Mb)", "STATUS", "VMs", "TYPE")
					headerPrinted = true
				}
				fmt.Printf("%-45s\t%-10d\t%-10s\t%d\t%-10s\t\n", d.Name, d.SizeMb, d.Status, d.AttachedVmCount, d.BusTypeDesc)
			}
		}
	}
}
