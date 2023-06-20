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
	"os"
	"text/tabwriter"

	"github.com/giantswarm/cloud-director-cli/pkg/vcd/utils"
	"gopkg.in/yaml.v3"
)

const (
	config_env   = "CDCLI_CONFIG"
	home_env     = "HOME"
	cfg_dirname  = ".cd-cli"
	cfg_filename = "config.yaml"
)

type VCDInfo struct {
	Name         string `yaml:"name"`
	RefreshToken string `yaml:"refreshToken,omitempty"`
	Org          string `yaml:"org,omitempty"`
	Site         string `yaml:"site"`
	Ovdc         string `yaml:"ovdc,omitempty"`
	Insecure     bool   `yaml:"insecure,omitempty"`
	Username     string `yaml:"username,omitempty"`
	Password     string `yaml:"password,omitempty"`
}

type Config struct {
	Contexts       []VCDInfo `yaml:"contexts"`
	CurrentContext string    `yaml:"currentContext,omitempty"`
}

// it assumes the config to be placed in ~/.cd-cli/config.yaml
// or under path contained in $CDCLI_CONFIG env var
func getConfigPath() string {
	var path string
	path, found := os.LookupEnv(config_env)
	if !found {
		p, foundHome := os.LookupEnv(home_env)
		path = fmt.Sprintf("%s/%s/%s", p, cfg_dirname, cfg_filename)
		if !foundHome {
			log.Fatal(fmt.Sprintf("Place your config to $%s/%s/%s or set %s", home_env, cfg_dirname, cfg_filename, config_env))
		}
	}
	return path
}

func parseConfig() (*Config, error) {
	path := getConfigPath()

	yfile, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(fmt.Errorf("Unable to open configuration file, make sure it exists (%s)\n%w", path, err))
	}
	data := &Config{}
	err2 := yaml.Unmarshal(yfile, &data)

	if err2 != nil {
		log.Fatal(err2)
	}
	if len(data.Contexts) == 0 {
		log.Fatal(fmt.Errorf("No contexts were specified in (%s)", path))
	}

	return data, nil
}

func persistConfig(cfg *Config) {
	path := getConfigPath()
	bytes, err := yaml.Marshal(cfg)
	if err != nil {
		log.Fatal(fmt.Errorf("Marshalling failed: \n%w", err))
	}
	if err2 := os.WriteFile(path, bytes, 644); err2 != nil {
		log.Fatal(fmt.Errorf("Unable to persist the config back to the file on path (%s)\n%w", path, err2))
	}
}

func ChangeCurrentContext(context string) {
	cfg, err := parseConfig()
	if err != nil {
		log.Fatal(fmt.Errorf("Failed to parse the config: \n%w", err))
	}
	contextFound := utils.Contains(utils.Map(cfg.Contexts, func(i VCDInfo) string {
		return i.Name
	}), context)
	if !contextFound {
		log.Fatal(fmt.Sprintf("Unable to switch to non-existent context: '%s'", context))
	}
	cfg.CurrentContext = context
	persistConfig(cfg)
}

func PrintCurrentContext() {
	cfg, err := parseConfig()
	if err != nil {
		log.Fatal(fmt.Errorf("Failed to parse the config: \n%w", err))
	}
	availableContexts := utils.Map(cfg.Contexts, func(i VCDInfo) []string {
		return []string{i.Name, i.Site}
	})
	writer := tabwriter.NewWriter(os.Stdout, 0, 8, 4, '\t', 0)
	utils.Printf(writer, "%s\t%s\t%s\t", "CURRENT", "NAME", "SITE")
	utils.Printf(writer, "\n")
	var isCurrent rune
	for _, ctx := range availableContexts {
		isCurrent = ' '
		if ctx[0] == cfg.CurrentContext {
			isCurrent = '*'
		}
		utils.Printf(writer, "%c\t%s\t%s\t", isCurrent, ctx[0], ctx[1])
		utils.Printf(writer, "\n")
	}
	err = writer.Flush()
	if err != nil {
		log.Fatal(err)
	}
}
