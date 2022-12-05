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
	delVmsCmd     = &cobra.Command{
		Use:   "vms",
		Short: "Delete the VMs",
		Long:  `Delete the VMs in the cloud director`,
		Run: func(cmd *cobra.Command, args []string) {
			switch cmd.Parent().Name() {
			case "list":
				vcd.PrintVMs(verbose, onlyTemplates, "")
			}
		},
	}
)

func init() {
	cleanCmd.AddCommand(delVmsCmd)
	//delVmsCmd.Flags().BoolVarP(&onlyTemplates, "onlyTemplates", "t", false, "List only templates")
}
