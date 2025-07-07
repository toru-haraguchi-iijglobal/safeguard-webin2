// functions to handle chromedp for webin2.

package main

import (
	"log"

	"github.com/chromedp/chromedp"
)

// build actions array
func build_actions(json_line JsonLine) []chromedp.Action {
	actions := []chromedp.Action{}
	log.Println(json_line)
	return actions
}

func run(actions []chromedp.Action) bool {
	log.Printf("Running chromedp with %d actions...", len(actions))
	return true
}
