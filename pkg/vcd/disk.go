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
	"github.com/vmware/go-vcloud-director/v2/govcd"
	"github.com/vmware/go-vcloud-director/v2/types/v56"

	"github.com/giantswarm/cloud-director-cli/pkg/vcd/utils"
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

	results := make([]*types.DiskRecordType, 0)
	page := 1
	for {
		notEncodedParams := map[string]string{"type": "disk", "filter": filter, "filterEncoded": "true", "page": strconv.Itoa(page)}
		pageResult, err := manager.Client.VDC.QueryWithNotEncodedParams(nil, notEncodedParams)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, pageResult.Results.DiskRecord...)
		page++

		if !utils.HasNextPageLink(pageResult.Results.Link) {
			break
		}
	}

	return results
}

func (manager *DiskManager) Delete(names []string, detach bool) {
	for _, name := range names {
		fmt.Printf("Deleting disk:[%s]\n", name)
		disks, err := manager.Client.VDC.GetDisksByName(name, false)
		if err != nil {
			log.Fatal(err)
		}
		for _, d := range *disks {
			if detach {
				e := manager.detachFromAllVMs(&d)
				if e != nil {
					log.Fatal(e)
				}
			}

			t, e := d.Delete()
			if e != nil {
				log.Fatal(e)
			}
			fmt.Printf("delete task was created: %s\n", t.Task.ID)
		}
	}
}

func (manager *DiskManager) detachFromAllVMs(d *govcd.Disk) error {
	params := &types.DiskAttachOrDetachParams{
		Disk: &types.Reference{HREF: d.Disk.HREF},
	}

	vmHrefs, err := d.GetAttachedVmsHrefs()
	if err != nil {
		return fmt.Errorf("unable to get attached vms from disk: [%s]", d.Disk.Name)
	}

	for _, vmHref := range vmHrefs {
		vm, err := manager.Client.VCDClient.Client.GetVMByHref(vmHref)
		if err != nil {
			return fmt.Errorf("unable to get vm by href: [%s]", vmHref)
		}

		fmt.Printf("Detaching disk:[%s] from vm:[%s]\n", d.Disk.Name, vm.VM.Name)
		task, err := vm.DetachDisk(params)
		if err != nil {
			return fmt.Errorf("unable to create task to detack disk: [%v]", err)
		}

		err = task.WaitTaskCompletion()
		if err != nil {
			return fmt.Errorf("unable to detack disk: [%v]", err)
		}

	}

	// refresh before deleting to fetch "remove" links
	return d.Refresh()
}
