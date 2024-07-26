package utils

import "github.com/vmware/go-vcloud-director/v2/types/v56"

func HasNextPageLink(link []*types.Link) bool {
	for _, l := range link {
		if l.Rel == "nextPage" {
			return true
		}
	}
	return false
}
