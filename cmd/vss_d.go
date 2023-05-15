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
)

var (
	failifabsent bool
	delVssCmd    = &cobra.Command{
		Use:     "vss VSname... ",
		Aliases: []string{"vs", "virtualservices", "virtualservice", "virtualsvc", "virtualsvcs", "vsvc", "vsvcs"},
		Short:   "Delete the virtual service",
		Long: `Delete the virtual service in the cloud director

	Example:
	--------
	cd-cli delete vs sdf -y --failifabsent
`,
		Run: func(cmd *cobra.Command, args []string) {
			manager := vcd.VirtualServiceManager{
				Client: vcdClient,
			}
			manager.Delete(args, failifabsent, network)
		},
	}
)

func init() {
	cleanCmd.AddCommand(delVssCmd)
	delVssCmd.Flags().BoolVar(&failifabsent, "failifabsent", false, "command will return non-zero code if the virtual service is not there")
	addNetworkFlag(delVssCmd)
}
