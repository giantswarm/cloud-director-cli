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

	"github.com/spf13/cobra"

	"github.com/giantswarm/cloud-director-cli/pkg/vcd"
	"github.com/giantswarm/cloud-director-cli/pkg/vcd/utils"
)

var (
	href            string
	listMetadataCmd = &cobra.Command{
		Use:   "metadata",
		Short: "List metadata of a resources such as vapp",
		Long: `List metadata of a resources such as vapp

	Example:
	--------
	cd-cli list metadata -href="https://mydomain.com/api/vApp/vapp-c7b89940-eece-48c7-8895-883fd347ee3e"
`,
		Run: func(cmd *cobra.Command, args []string) {
			manager := vcd.MetadataManager{
				Client: vcdClient,
			}
			items := manager.List(href)
			utils.Print(outputFormat, items, "Key",
				[]string{"KEY", "Value"},
				[]string{"Key", "TypedValue.Value"})
		},
	}
)

func init() {
	listCmd.AddCommand(listMetadataCmd)
	listMetadataCmd.Flags().StringVar(&href, "href", "", "HREF of a cloud-director resource")
	err := listMetadataCmd.MarkFlagRequired("href")
	if err != nil {
		log.Fatal(err)
	}
}
