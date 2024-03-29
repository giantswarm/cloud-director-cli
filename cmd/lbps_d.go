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
	cascade    bool
	delLbpsCmd = &cobra.Command{
		Use:     "lbps LBPoolName...",
		Aliases: []string{"lbp", "lbpool", "lbpools"},
		Short:   "Delete the LBPool",
		Long: `Delete the Load Balancer Pool in the cloud director

	Example:
	--------
	cd-cli delete lbp sdf -y --failifabsent --cascade
`,
		Run: func(cmd *cobra.Command, args []string) {
			manager := vcd.LoadBalancerPoolManager{
				Client: vcdClient,
			}
			manager.Delete(args, failifabsent, verbose, cascade, network)
		},
	}
)

func init() {
	cleanCmd.AddCommand(delLbpsCmd)
	delLbpsCmd.Flags().BoolVar(&cascade, "cascade", false, "delete also the associated virual services, this assumes them to have the same name as the LB pools")
	delLbpsCmd.Flags().BoolVar(&failifabsent, "failifabsent", false, "command will return non-zero code if the load balancer pool is not there")
	addNetworkFlag(delLbpsCmd)
}
