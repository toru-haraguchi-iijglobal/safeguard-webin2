// functions to handle JSONL file for webin2.
package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

func search_jsonlines(asset string, jsonl_filename string) Definition {
	// Open JSON Lines file
	jsonlfile, err := os.Open(jsonl_filename)
	if err != nil {
		log.Fatalf("File \"%s\" could not be opened.", jsonl_filename)
	}
	defer jsonlfile.Close()
	decoder := json.NewDecoder(jsonlfile)

	// Scan jsonlfile for the asset specified
	for {
		var json_line Definition
		err := decoder.Decode(&json_line)
		if err != nil {
			if err == io.EOF {
				log.Printf("Asset \"%s\" not found", asset)
				return Definition{}
			} else {
				log.Println("JSON Line contains error. Last correct asset: " + json_line.Asset)
				log.Fatal(err)
			}
		} else if json_line.Asset == asset {
			log.Printf("Asset \"%s\" found", asset)
			// If match, return the JsonLine
			return json_line
		}
	}

}
