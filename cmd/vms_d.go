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
	"github.com/giantswarm/cloud-director-cli/pkg/vcd"
	"github.com/spf13/cobra"
)

var (
	onlyTemplates bool
	yes           bool
	delVmsCmd     = &cobra.Command{
		Use:     "vms -a vAppName VMname... ",
		Aliases: []string{"vm"},
		Short:   "Delete the VMs",
		Long: `Delete the VMs in the cloud director of a given vApp

	Example:
	--------
	cd-cli clean vms --vapp=jiri3 jiri3-worker-7b4d46494-8rj59 jiri3-worker-7b4d46494-p6vhp
	Are you sure you want to delete following VMs: [jiri3-worker-7b4d46494-8rj59, jiri3-worker-7b4d46494-p6vhp] [y/n]?
	y
	
	or for non-interactive mode:
	cd-cli clean vms -y --vapp=jiri3 jiri3-worker-7b4d46494-8rj59 jiri3-worker-7b4d46494-p6vhp
`,
		Run: func(cmd *cobra.Command, args []string) {
			vcd.DeleteVMs(args, vapp, yes, verboseClient)
		},
	}
)

func init() {
	cleanCmd.AddCommand(delVmsCmd)
	delVmsCmd.Flags().StringVarP(&vapp, "vapp", "a", "", "vApp whose VMs will be deleted")
	delVmsCmd.MarkFlagRequired("vapp")
	delVmsCmd.Flags().BoolVarP(&yes, "assumeyes", "y", false, "non-interactive mode assuming yes to all questions")
}
