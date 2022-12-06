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
	"github.com/vmware/cloud-provider-for-cloud-director/pkg/vcdsdk"
	"github.com/vmware/go-vcloud-director/v2/types/v56"
	"log"
)

func ListvApps(verboseClient bool) []*types.ResourceReference {
	cache := Cache{}
	c, e := cache.CachedClient(verboseClient)
	if e != nil {
		log.Fatal(e)
	}
	vapps := c.VDC.GetVappList()
	return vapps
}

func DeletevApp(name string, yes bool, verboseClient bool) error {
	cache := Cache{}
	c, e := cache.CachedClient(verboseClient)
	if e != nil {
		log.Fatal(e)
	}
	m, err := vcdsdk.NewVDCManager(c, "", "")
	if err != nil {
		log.Fatal(err)
	}
	if !yes {
		fmt.Printf("Are you sure you want to delete vApp '%s'[y/n]?\n", name)
		var char rune
		_, err := fmt.Scanf("%c", &char)
		if err != nil {
			log.Fatal(err)
		}
		if char != 'y' && char != 'Y' {
			return nil
		}
	}
	err2 := m.DeleteVApp(name)
	if err2 != nil {
		log.Fatal(err2)
	}
	return nil
}

func PrintvApps(verbose bool, verboseClient bool, onlyTemplates bool, vapp string) error {
	var headerPrinted bool
	for _, vapp := range ListvApps(verboseClient) {
		if !verbose {
			fmt.Println(vapp.Name)
		} else {
			if !headerPrinted {
				fmt.Printf("%-35s\t%-16s\t\n", "NAME", "ID")
				headerPrinted = true
			}
			fmt.Printf("%-35s\t%-16s\t\n", vapp.Name, vapp.ID)
		}
	}
	return nil
}
