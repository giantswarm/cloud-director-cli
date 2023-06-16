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
	"log"

	"github.com/giantswarm/cloud-director-cli/pkg/vcd/client"
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var (
	configCmd = &cobra.Command{
		Use:   "config",
		Short: "Switch current context, list available contexts.",
		Long: `This command can list all available contexts each representing different VCD, or set the currently active one

	Examples:
	---------
	cd-cli config get-contexts
	cd-cli config set-context foo
`,
	}
	getContexts = &cobra.Command{
		Use:   "get-contexts",
		Short: "list available contexts.",
		Run: func(cmd *cobra.Command, args []string) {
			client.PrintCurrentContext()
		},
	}
	setContext = &cobra.Command{
		Use:   "set-context",
		Short: "set current default context.",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				log.Fatal("cd-cli config set-context takes exactly only one argument with the new desired default context name.")
				return
			}
			client.ChangeCurrentContext(args[0])
		},
	}
)

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(getContexts)
	configCmd.AddCommand(setContext)
}
