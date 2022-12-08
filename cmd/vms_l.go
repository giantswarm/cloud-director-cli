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

// listVmsCmd represents the vms command
var (
	vapp       string
	listVmsCmd = &cobra.Command{
		Use:     "vms",
		Aliases: []string{"vm"},
		Short:   "List all the VMs",
		Long: `List all the VMs in the cloud director (or vm templates)

	Examples:
	--------
	cd-cli list vms -v
	NAME                               	VAPP            	STATUS    	DEPLOYED
	guppy-8fb68                        	guppy           	POWERED_ON	true
	guppy-w4chm                        	guppy           	POWERED_ON	true
	guppy-worker-79fbbb5b7c-9mvpm      	guppy           	POWERED_ON	true
	squid-proxy                        	installation-proxy	POWERED_ON	true

	cd-cli list vms --vapp installation-proxy
	squid-proxy
`,
		Run: func(cmd *cobra.Command, args []string) {
			vcd.PrintVMs(verbose, verboseClient, onlyTemplates, vapp)
		},
	}
)

func init() {
	listCmd.AddCommand(listVmsCmd)
	listVmsCmd.Flags().StringVarP(&vapp, "vapp", "a", "", "Only VMs/templates of this vAPP will be listed")
	listVmsCmd.Flags().BoolVarP(&onlyTemplates, "onlyTemplates", "t", false, "List only templates")
}
