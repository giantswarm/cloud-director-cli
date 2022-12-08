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
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// listCmd represents the list command
var (
	output  string
	listCmd = &cobra.Command{
		Use:   "list",
		Short: "This command can list various resources from vcd: vms, vapps, disks, virtual services, etc.",
		Long: `This command can list various resources from vcd: vms, vapps, disks, virtual services, etc.

	Examples: 
	---------
	cd-cli list vms -v
	cd-cli list disks
`,
		PreRun: ValidateOutput,
	}
)

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.PersistentFlags().StringVarP(&output, "output", "o", "names", "Output format. One of: (json, yaml, names, columns)")
}

func ValidateOutput(cmd *cobra.Command, _ []string) {
	switch cmd.Flag("output").Value.String() {
	case "json", "yaml", "names", "columns": //no-op
	default:
		fmt.Printf("\nunknown value for --output '%s'\n", cmd.Flag("output").Value.String())
		fmt.Printf("use one of the following: [json, yaml, names, columns]\n")
		os.Exit(1)
	}
}
