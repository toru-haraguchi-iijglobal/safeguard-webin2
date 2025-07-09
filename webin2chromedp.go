// functions to handle chromedp for webin2.

package main

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

func run(json_line Definition) bool {
	log.Printf("Running chromedp with %d actions...", len(json_line.Actions))

	// Setting up browser options
	opts := append(
		chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
		chromedp.Flag("enable-automation", false),
		chromedp.Flag("hide-scrollbars", false),
		chromedp.Flag("mute-audio", false),
		chromedp.Flag("disable-infobars", true),
		chromedp.Flag("window-size", "1280,800"),
	)
	if json_line.UseEdge {
		opts = append(opts,
			chromedp.ExecPath("C:\\Program Files (x86)\\Microsoft\\Edge\\Application\\msedge.exe"),
		)
	}
	if json_line.Secret {
		opts = append(opts,
			chromedp.Flag("incognito", true),
		)
	}
	if !json_line.CertValidation {
		opts = append(opts,
			chromedp.Flag("ignore-certificate-errors", true),
		)
	}

	// Initialize contexts
	allocCtx, _ := chromedp.NewExecAllocator(context.Background(), opts...)
	runCtx, _ := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))

	// Build actions based on the JSON line
	actions := []chromedp.Action{}
	for i, action := range json_line.Actions {
		log.Printf("%d %s %s %d", i,
			action.Type,
			action.Target,
			action.Value)

		switch action.Type {
		case "navigate":
			actions = append(actions,
				chromedp.Navigate(action.Target))
		case "click":
			actions = append(actions,
				chromedp.Click(action.Target,
					chromedp.ByQuery,
					chromedp.NodeVisible))
		case "account":
			actions = append(actions,
				chromedp.SendKeys(action.Target,
					account,
					chromedp.ByQuery,
					chromedp.NodeVisible))
		case "password":
			actions = append(actions,
				chromedp.SendKeys(action.Target,
					password,
					chromedp.ByQuery,
					chromedp.NodeVisible))
		case "sleep":
			actions = append(actions,
				chromedp.Sleep(time.Millisecond*time.Duration(action.Value)))

		default:
		}
	}

	err := chromedp.Run(runCtx, actions...)

	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("Done")
	}

	return true
}
