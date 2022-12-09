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
	listVappsCmd = &cobra.Command{
		Use:     "vapps",
		Aliases: []string{"vapp", "virtualapp", "virtualapps"},
		Short:   "List all the vApps",
		Long: `List all the vApps in the cloud director

	Example:
	--------
	cd-cli list vapp -v
	NAME                               	ID
	guppy                              	urn:vcloud:vapp:afe1a37f-4b7d-4c0f-a5f3-14f19bf5f073
	installation-proxy                 	urn:vcloud:vapp:8994a22f-4870-43d4-8897-6945f2e96d9b
	gs-eric-vcd                        	urn:vcloud:vapp:26f79f84-908b-4ee8-88a9-36d5066175f8
`,
		Run: func(cmd *cobra.Command, args []string) {
			vcd.PrintvApps(output, verbose)
		},
		PreRun: ValidateOutput,
	}
)

func init() {
	listCmd.AddCommand(listVappsCmd)
}
