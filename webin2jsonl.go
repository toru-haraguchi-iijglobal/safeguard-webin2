// functions to handle JSONL file for webin2.
package main

import (
	"encoding/json"
	"io"
	"os"
)

func search_jsonlines(asset string, jsonl_filename string) Definition {
	LogAssetSearch(asset, jsonl_filename)
	LogFileOperation("Opening", jsonl_filename)

	// Open JSON Lines file
	jsonlfile, err := os.Open(jsonl_filename)
	if err != nil {
		Fatal("File \"%s\" could not be opened: %v", jsonl_filename, err)
	}
	defer jsonlfile.Close()
	decoder := json.NewDecoder(jsonlfile)

	// Scan jsonlfile for the asset specified
	lineNumber := 0
	for {
		var json_line Definition
		err := decoder.Decode(&json_line)
		lineNumber++

		if err != nil {
			if err == io.EOF {
				LogAssetNotFound(asset, jsonl_filename)
				return Definition{}
			} else {
				Error("JSON Line %d contains error. Last correct asset: %s", lineNumber, json_line.Asset)
				Fatal("Failed to parse JSON Line: %v", err)
			}
		} else if json_line.Asset == asset {
			LogAssetFound(asset)
			Debug("Asset found at line %d in %s", lineNumber, jsonl_filename)
			// If match, return the JsonLine
			return json_line
		} else {
			Debug("Line %d: Asset '%s' does not match, continuing search", lineNumber, json_line.Asset)
		}
	}
}
