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
	"fmt"
	"log"
	"net/url"
	"os"
	"sync"

	"github.com/vmware/cloud-provider-for-cloud-director/pkg/vcdsdk"
	"gopkg.in/yaml.v3"
)

const (
	config_env   = "CDCLI_CONFIG"
	home_env     = "HOME"
	cfg_dirname  = ".cd-cli"
	cfg_filename = "config.yaml"
)

type Config struct {
	RefreshToken string `yaml:"refreshToken,omitempty"`
	Org          string `yaml:"org,omitempty"`
	Site         string `yaml:"site"`
	Ovdc         string `yaml:"ovdc,omitempty"`
	Insecure     bool   `yaml:"insecure,omitempty"`
	Username     string `yaml:"username,omitempty"`
	Password     string `yaml:"password,omitempty"`
}

type Cache struct {
	atStart sync.Once
	client  *vcdsdk.Client
}

func (cache *Cache) lazyInit(items bool) {
	cache.atStart.Do(func() {
		cfg, err := parseConfig()
		if err != nil {
			log.Fatal(err)
		}
		cache.client, err = makeClient(cfg, items)
		if err != nil {
			log.Fatal(err)
		}
	})
}

// it assumes the config to be placed in ~/.cd-cli/config.yaml
// or under path contained in $CDCLI_CONFIG env var
func parseConfig() (*Config, error) {
	var path string
	path, found := os.LookupEnv(config_env)
	if !found {
		p, foundHome := os.LookupEnv(home_env)
		path = fmt.Sprintf("%s/%s/%s", p, cfg_dirname, cfg_filename)
		if !foundHome {
			log.Fatal(fmt.Sprintf("Place your config to $%s/%s/%s or set %s", home_env, cfg_dirname, cfg_filename, config_env))
		}
	}
	yfile, err := os.ReadFile(fmt.Sprintf("%s", path))
	if err != nil {
		log.Fatal(fmt.Errorf("Unable to open configuration file, make sure it exist (~/.cd-cli/config.yaml)\n%w", err))
	}
	data := &Config{}
	err2 := yaml.Unmarshal(yfile, &data)

	if err2 != nil {
		log.Fatal(err2)
	}

	return data, nil
}

func makeClient(cfg *Config, items bool) (*vcdsdk.Client, error) {
	_, err := url.ParseRequestURI(cfg.Site)
	if err != nil {
		return nil, fmt.Errorf("unable to parse site url: %s", err)
	}
	if !items { // the underlying vcd client prints a lot of messages to stderr
		os.Stderr = nil
	}
	client, err := vcdsdk.NewVCDClientFromSecrets(cfg.Site, cfg.Org,
		cfg.Ovdc, cfg.Org, cfg.Username, cfg.Password, cfg.RefreshToken, cfg.Insecure, true)
	if err != nil {
		return nil, fmt.Errorf("unable to authenticate: %s", err)
	}

	return client, nil
}

func (cache *Cache) CachedClient(items bool) (*vcdsdk.Client, error) {
	cache.lazyInit(items)
	return cache.client, nil
}
