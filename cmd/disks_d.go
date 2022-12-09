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
	delDisksCmd = &cobra.Command{
		Use:     "disks diskname...",
		Aliases: []string{"disk"},
		Short:   "Delete the disk",
		Long: `Delete the disk in the cloud director

	Example:
	--------
	cd-cli delete disks sdf1 sdf2 -y
`,
		Run: func(cmd *cobra.Command, args []string) {
			vcd.DeleteDisks(args, yes, verbose)
		},
	}
)

func init() {
	cleanCmd.AddCommand(delDisksCmd)
	delDisksCmd.Flags().BoolVar(&failifabsent, "failifabsent", false, "command will return non-zero code if the load balancer pool is not there")
}
