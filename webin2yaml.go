// functions to handle JSONL file for webin2.
package main

import (
	"io"
	"os"

	"gopkg.in/yaml.v2"
)

func search_yaml(asset string, yaml_filename string) Definition {
	LogAssetSearch(asset, yaml_filename)
	LogFileOperation("Opening", yaml_filename)

	// Open YAML file
	yamlfile, err := os.Open(yaml_filename)
	if err != nil {
		Fatal("File \"%s\" could not be opened: %v", yaml_filename, err)
	}
	defer yamlfile.Close()
	decoder := yaml.NewDecoder(yamlfile)

	// Scan yamlfile for the asset specified
	documentNumber := 0
	for {
		var yaml Definition
		err := decoder.Decode(&yaml)
		documentNumber++

		if err != nil {
			if err == io.EOF {
				LogAssetNotFound(asset, yaml_filename)
				return Definition{}
			} else {
				Error("YAML document %d contains error. Last correct asset: %s", documentNumber, yaml.Asset)
				Fatal("Failed to parse YAML: %v", err)
			}
		} else if yaml.Asset == asset {
			LogAssetFound(asset)
			Debug("Asset found at document %d in %s", documentNumber, yaml_filename)
			// If match, return the YAML definition
			return yaml
		} else {
			Debug("Document %d: Asset '%s' does not match, continuing search", documentNumber, yaml.Asset)
		}
	}
}
