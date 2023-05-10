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
	listLbpsCmd = &cobra.Command{
		Use:     "lbps",
		Aliases: []string{"lbp", "lbpool", "lbpools"},
		Short:   "List all the LB Pools",
		Long: `List all the Load Balancer Pools in the cloud director

	Examples:
	--------
	cd-cli list lbps -o=columns
	NAME                                                                                      	ALGOTITHM        	MEMBERS
	ingress-pool-nginx-ingress-controller-app--http                                           	LEAST_CONNECTIONS	6
	ingress-pool-nginx-ingress-controller-app--https                                          	LEAST_CONNECTIONS	6
	gs-eric-vcd-NO_RDE_b03a4df5-585f-48a9-8916-d378c44b7c16-tcp                               	ROUND_ROBIN      	1
	ingress-pool-nginx-ingress-controller-app-NO_RDE_b03a4df5-585f-48a9-8916-d378c44b7c16-http	LEAST_CONNECTIONS	4
	ingress-pool-nginx-ingress-controller-app-NO_RDE_b03a4df5-585f-48a9-8916-d378c44b7c16-https	LEAST_CONNECTIONS	4
	ingress-pool-nginx-ingress-controller-app-NO_RDE_ca501275-f986-4d50-a6ec-e084341d15d2-http	LEAST_CONNECTIONS	6
	ingress-pool-nginx-ingress-controller-app-NO_RDE_ca501275-f986-4d50-a6ec-e084341d15d2-https	LEAST_CONNECTIONS	6
	guppy-NO_RDE_ca501275-f986-4d50-a6ec-e084341d15d2-tcp                                     	ROUND_ROBIN      	3
`,
		Run: func(cmd *cobra.Command, args []string) {
			manager := vcd.LoadBalancerPoolManager{
				Client: vcdClient,
			}
			items := manager.List(network)
			manager.Print(output, items)
		},
	}
)

func init() {
	listCmd.AddCommand(listLbpsCmd)
	addNetworkFlag(listLbpsCmd)
}
