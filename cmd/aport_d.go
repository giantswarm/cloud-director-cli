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
	delAportsCmd = &cobra.Command{
		Use:     "aport name... ",
		Aliases: []string{"aports", "appport", "appports"},
		Short:   "Delete the application port profiles",
		Long: `Delete the application port profiles in the cloud director

	Example:
	--------
	cd-cli delete aport sdf -y --failifabsent
`,
		Run: func(cmd *cobra.Command, args []string) {
			vcd.DeleteAport(args, failifabsent, verbose, network)
		},
	}
)

func init() {
	cleanCmd.AddCommand(delAportsCmd)
	delAportsCmd.Flags().BoolVar(&failifabsent, "failifabsent", false, "command will return non-zero code if the application port profile is not there")
	addNetworkFlag(delAportsCmd)
}
