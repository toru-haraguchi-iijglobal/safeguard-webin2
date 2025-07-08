// functions to handle JSONL file for webin2.
package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type JsonLine struct {
	Asset          string   `json:"asset"`           // -asset=[asset name of SPP]
	UseEdge        bool     `json:"use_edge"`        // Use MS Edge instead of Google Chrome
	Secret         bool     `json:"secret"`          // Open with Secret Window
	CertValidation bool     `json:"cert_validation"` // Do or Don't Cert Validation
	Actions        []Action `json:"actions"`         // array of actions to perform
}

type Action struct {
	Type   string `json:"type"`   // chromedp action type
	Target string `json:"target"` // url or element selector
	Value  int    `json:"value"`  // value to set or click
}

func search_jsonlines(asset string, jsonl_filename string) JsonLine {
	// Open JSON Lines file
	jsonlfile, err := os.Open(jsonl_filename)
	if err != nil {
		log.Fatalln("'" + jsonl_filename + "' file could not be opened.")
	}
	defer jsonlfile.Close()
	decoder := json.NewDecoder(jsonlfile)

	// Scan jsonlfile for the asset specified
	for {
		var json_line JsonLine
		err := decoder.Decode(&json_line)
		if err != nil {
			if err == io.EOF {
				log.Printf("Asset \"%s\" not found", asset)
				return JsonLine{}
			} else {
				log.Println("Last correct asset: " + json_line.Asset)
				log.Fatal(err)
			}
		} else {
			if json_line.Asset == asset {
				log.Printf("Asset \"%s\" found", asset)
				// If match, return the JsonLine
				return json_line
			}
		}
	}

}
