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
)

// cleanCmd represents the clean command
var cleanCmd = &cobra.Command{
	Use:     "clean",
	Aliases: []string{"delete"},
	Short:   "This command can clean various resources from vcd: vms, vapps, virtual services, etc.",
	Long: `This command can clean various resources from vcd: vms, vapps, virtual services, etc.

	Examples: 
	---------
	cd-cli clean vapp jiri3
	Are you sure you want to delete vApp 'jiri3'[y/n]?
	y
	
	cd-cli clean vms --assumeyes --vapp=jiri3 jiri3-worker-7b4d46494-8rj59 jiri3-worker-7b4d46494-p6vhp

	cd-cli clean virtualservice --assumeyes guppy-NO_RDE_ca501275-f986-4d50-a6ec-e084341d15d2-tcp
`,
}

func init() {
	rootCmd.AddCommand(cleanCmd)
	cleanCmd.PersistentFlags().BoolVarP(&yes, "assumeyes", "y", false, "non-interactive mode assuming yes to all questions")
}
