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

package vcd

import (
	"log"
	"net/http"

	"github.com/vmware/cloud-provider-for-cloud-director/pkg/vcdsdk"
	"github.com/vmware/go-vcloud-director/v2/types/v56"
)

type MetadataManager struct {
	Client *vcdsdk.Client
}

func (manager *MetadataManager) List(href string) []*types.MetadataEntry {
	metadata := &types.Metadata{}
	_, err := manager.Client.VCDClient.Client.ExecuteRequest(href+"/metadata/", http.MethodGet, types.MimeMetaData, "error retrieving metadata: %s", nil, metadata)
	if err != nil {
		log.Fatal(err)
	}
	// hide for cleaner output
	for _, item := range metadata.MetadataEntry {
		item.Link = nil
	}
	return metadata.MetadataEntry
}
