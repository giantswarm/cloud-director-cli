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

package cmd

import (
	"github.com/spf13/cobra"

	"github.com/giantswarm/cloud-director-cli/pkg/vcd"
	"github.com/giantswarm/cloud-director-cli/pkg/vcd/utils"
)

var (
	unattached   bool
	listDisksCmd = &cobra.Command{
		Use:     "disks",
		Aliases: []string{"disk"},
		Short:   "List all the disks",
		Long: `List all the disks in the cloud director

	Example:
	--------
	cd-cli list disks -o=columns
	NAME                                         	SIZE(Mb)  	STATUS    	VMs	TYPE
	pvc-69969a35-b9df-4605-b052-d60beabf0d20     	5120      	RESOLVED  	0	Paravirtual (SCSI)
	pvc-37eef8f3-8708-40fb-b4c3-6d6cc3e0a760     	1024      	RESOLVED  	0	Paravirtual (SCSI)
	pvc-5add9939-513c-4017-a76b-927221881ac1     	1024      	RESOLVED  	0	Paravirtual (SCSI)
	...
`,
		Run: func(cmd *cobra.Command, args []string) {
			manager := vcd.DiskManager{
				Client: vcdClient,
			}

			items := manager.List(vcd.DiskListParams{Unattached: unattached})

			utils.Print(outputFormat, items, "Name",
				[]string{"NAME", "SIZE(Mb)", "STATUS", "VMs", "TYPE"},
				[]string{"Name", "SizeMb", "Status", "AttachedVmCount", "BusTypeDesc"})
		},
	}
)

func init() {
	listCmd.AddCommand(listDisksCmd)
	listDisksCmd.Flags().BoolVarP(&unattached, "unattached", "u", false, "Only unattached disks will be listed")
}
