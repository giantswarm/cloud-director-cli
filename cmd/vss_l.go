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
	listVssCmd = &cobra.Command{
		Use:     "vss",
		Aliases: []string{"vs", "virtualservices", "virtualservice", "virtualsvc", "virtualsvcs", "vsvc", "vsvcs"},
		Short:   "List all the virtual services",
		Long: `List all the virtual services in the cloud director

	Examples:
	--------
	cd-cli list vs -o=columns
	NAME                                                                                        IP                  HEALTH
	gs-eric-vcd-NO_RDE_b03a4df5-585f-48a9-8916-d378c44b7c16-tcp                                 178.170.32.55       UP
	ingress-vs-ingress-nginx-controller-NO_RDE_b03a4df5-585f-48a9-8916-d378c44b7c16-http        192.168.8.6         UP
	ingress-vs-ingress-nginx-controller-NO_RDE_b03a4df5-585f-48a9-8916-d378c44b7c16-https       192.168.8.7         UP
	ingress-vs-ingress-nginx-controller--http                                                   192.168.8.4         UP
`,
		Run: func(cmd *cobra.Command, args []string) {
			manager := vcd.VirtualServiceManager{
				Client: vcdClient,
			}
			items := manager.List(vcd.VirtualServiceListParams{Network: network})
			utils.Print(outputFormat, items, "name",
				[]string{"NAME", "IP", "HEALTH"},
				[]string{"name", "virtualIpAddress", "healthStatus"})
		},
	}
)

func init() {
	listCmd.AddCommand(listVssCmd)
	addNetworkFlag(listVssCmd)
}
