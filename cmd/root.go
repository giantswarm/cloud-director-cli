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
	"os"
	"strings"

	"runtime/debug"

	"github.com/spf13/cobra"
)

var version string

// rootCmd represents the base command when called without any subcommands
var (
	verbose       bool
	verboseClient bool
	rootCmd       = &cobra.Command{
		Use:     "cd-cli",
		Version: getVersion(),
		Short:   "Simple cli tool that communicates with cloud director",
		Long: `cd-cli simple cli tool that communicates with cloud director

	Find more information at: https://github.com/giantswarm/cloud-director-cli
	
	Examples:
	---------
	cd-cli clean vms --assumeyes --vapp=jiri3 jiri3-worker-7b4d46494-8rj59 jiri3-worker-7b4d46494-p6vhp
	cd-cli clean vapp jiri3
	cd-cli list disks -v
`,
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cloud-director-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Use verbose output")
	rootCmd.PersistentFlags().BoolVarP(&verboseClient, "verboseClient", "d", false, "Use even more verbose output")

}

func getVersion() string {
	if version == "" {
		bi, ok := debug.ReadBuildInfo()
		if !ok {
			return ""
		}

		for _, dep := range bi.Deps {
			if strings.Contains(dep.Path, "cloud-director-cli") {
				version = dep.Version
			}
		}
	}
	return version
}
