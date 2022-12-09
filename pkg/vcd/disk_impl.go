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
	"github.com/vmware/go-vcloud-director/v2/types/v56"
	"log"
	"net/url"
)

func ListDisks(verbose bool) []*types.DiskRecordType {
	cache := Cache{}
	c, e := cache.CachedClient(verbose)
	if e != nil {
		log.Fatal(e)
	}
	filter := "vdc==" + url.QueryEscape(c.VDC.Vdc.HREF)
	notEncodedParams := map[string]string{"type": "disk", "filter": filter, "filterEncoded": "true"}
	results, err := c.VDC.QueryWithNotEncodedParams(nil, notEncodedParams)
	if err != nil {
		log.Fatal(err)
	}

	return results.Results.DiskRecord
}

func DeleteDisks(names []string, yes bool, verbose bool) {
	if len(names) == 0 {
		log.Fatal("Provide at least 1 name of a VM")
	}
	cache := Cache{}
	c, e := cache.CachedClient(verbose)
	if e != nil {
		log.Fatal(e)
	}
	if !yes {
		fmt.Printf("Are you sure you want to delete following disks: %v [y/n]?\n", names)
		var char rune
		_, err := fmt.Scanf("%c", &char)
		if err != nil {
			log.Fatal(err)
		}
		if char != 'y' && char != 'Y' {
			return
		}
	}
	for _, name := range names {
		disks, err := c.VDC.GetDisksByName(name, false)
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

func PrintDisks(output string, verbose bool, unattached bool) {
	items := ListDisks(verbose)
	switch output {
	case "json":
		PrintJson(items)
	case "yaml":
		PrintYaml(items)
	default:
		var headerPrinted bool
		for _, d := range items {
			if unattached && d.AttachedVmCount > 0 {
				continue
			}
			if output == "names" {
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
