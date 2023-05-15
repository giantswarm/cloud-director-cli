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

package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"text/tabwriter"

	"github.com/tidwall/gjson"
	"gopkg.in/yaml.v3"
)

func Print[T any](outputFormat string, items []T, namePath string, columnHeaders []string, columnPaths []string) {
	switch outputFormat {
	case "json":
		PrintJson(items)
	case "yaml":
		PrintYaml(items)
	case "names":
		PrintNames(items, namePath)
	default:
		PrintColumns(items, columnHeaders, columnPaths)
	}
}

func PrintNames[T any](items []T, namePath string) {
	for _, item := range items {
		fmt.Println(getValueByJsonPath(item, namePath))
	}
}

func getValueByJsonPath[T any](item T, path string) string {
	return gjson.Get(GetAsJson(item), path).String()
}

func PrintColumns[T any](items []T, columnHeaders []string, columnPaths []string) {
	var err error
	writer := tabwriter.NewWriter(os.Stdout, 0, 8, 4, '\t', 0)

	Printf(writer, "\n")
	for _, header := range columnHeaders {
		Printf(writer, "%s\t", header)
	}
	Printf(writer, "\n")

	for _, item := range items {
		for _, column := range columnPaths {
			Printf(writer, "%s\t", getValueByJsonPath(item, column))
		}
		Printf(writer, "\n")
	}

	err = writer.Flush()
	if err != nil {
		log.Fatal(err)
	}
}

func Printf(w io.Writer, format string, a ...any) {
	_, err := fmt.Fprintf(w, format, a...)
	if err != nil {
		log.Fatal(err)
	}
}

func PrintJson[T any](items []T) {
	j, err := json.Marshal(items)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(j))
}

func PrintYaml[T any](items []T) {
	y, err := yaml.Marshal(items)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(y))
}

func GetAsJson[T any](item T) string {
	j, err := json.Marshal(item)
	if err != nil {
		log.Fatal(err)
	}
	return string(j)
}
