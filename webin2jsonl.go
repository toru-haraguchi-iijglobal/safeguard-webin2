// functions to handle JSONL file for webin2.
package main

import "log"

type JsonLine struct {
	Asset   string   `json:"asset"`   // -asset=[asset name of SPP]
	Actions []Action `json:"actions"` // -actions=[array of actions to perform]
}

type Action struct {
	Type   string `json:"type"`   // chromedp action type
	Target string `json:"target"` // url or element selector
	Value  string `json:"value"`  // value to set or click
}

func search_jsonlines(jsonl_filename string) JsonLine {
	log.Println(jsonl_filename)
	var json_line JsonLine
	return json_line
}
