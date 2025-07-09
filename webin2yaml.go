// functions to handle JSONL file for webin2.
package main

import (
	"io"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

func search_yaml(asset string, yaml_filename string) Definition {
	// Open JSON Lines file
	yamlfile, err := os.Open(yaml_filename)
	if err != nil {
		log.Fatalf("File \"%s\" could not be opened.", yaml_filename)
	}
	defer yamlfile.Close()
	decoder := yaml.NewDecoder(yamlfile)

	// Scan yamlfile for the asset specified
	for {
		var yaml Definition
		err := decoder.Decode(&yaml)
		if err != nil {
			if err == io.EOF {
				log.Printf("Asset \"%s\" not found", asset)
				return Definition{}
			} else {
				log.Println("YAML contains error. Last correct asset: " + yaml.Asset)
				log.Fatal(err)
			}
		} else if yaml.Asset == asset {
			log.Printf("Asset \"%s\" found", asset)
			// If match, return the JsonLine
			return yaml
		}
	}

}
