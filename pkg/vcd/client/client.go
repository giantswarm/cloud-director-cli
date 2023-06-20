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

package client

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"sync"

	"github.com/vmware/cloud-provider-for-cloud-director/pkg/vcdsdk"
)

type Cache struct {
	atStart sync.Once
	client  *vcdsdk.Client
}

func (cache *Cache) lazyInit(verbose bool, context string) {
	cache.atStart.Do(func() {
		cfg, err := parseConfig()
		if err != nil {
			log.Fatal(err)
		}
		cache.client, err = makeClient(cfg, verbose, context)
		if err != nil {
			log.Fatal(err)
		}
	})
}

func makeClient(allContexts *Config, verbose bool, context string) (*vcdsdk.Client, error) {
	var cfg *VCDInfo
	currentContext := context
	if len(currentContext) == 0 {
		currentContext = allContexts.CurrentContext
	}
	if len(currentContext) == 0 {
		// current context is not specified, take the first occurence in the cfg file
		if len(allContexts.Contexts) > 0 {
			cfg = &allContexts.Contexts[0]
			currentContext = cfg.Name
		}
	} else {
		for _, c := range allContexts.Contexts {
			if c.Name == currentContext {
				cfg = &c
				break
			}
		}
		if cfg == nil {
			return nil, fmt.Errorf("Context with name '%s' was not found in the config file", currentContext)
		}
	}
	_, err := url.ParseRequestURI(cfg.Site)
	if err != nil {
		return nil, fmt.Errorf("unable to parse site url: %s", err)
	}
	if !verbose { // the underlying vcd client prints a lot of messages to stderr
		os.Stderr = nil
	} else {
		fmt.Printf("Using context: '%s'\nsite: '%s'\n", currentContext, cfg.Site)
	}
	client, err := vcdsdk.NewVCDClientFromSecrets(cfg.Site, cfg.Org,
		cfg.Ovdc, cfg.Org, cfg.Username, cfg.Password, cfg.RefreshToken, cfg.Insecure, true)
	if err != nil {
		return nil, fmt.Errorf("unable to authenticate: %s", err)
	}

	return client, nil
}

func (cache *Cache) CachedClient(verbose bool, context string) (*vcdsdk.Client, error) {
	cache.lazyInit(verbose, context)
	return cache.client, nil
}

func NewClient(verbose bool, context string) *vcdsdk.Client {
	cache := Cache{}
	c, e := cache.CachedClient(verbose, context)
	if e != nil {
		log.Fatal(e)
	}
	return c
}
