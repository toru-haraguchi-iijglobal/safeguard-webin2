// functions to handle chromedp for webin2.

package main

import (
	"context"
	"time"

	"github.com/chromedp/chromedp"
)

func run(definition Definition) bool {
	LogChromedpStart(len(definition.Actions))
	LogBrowserConfig(definition.UseEdge, definition.Secret, definition.CertValidation)

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
	if definition.UseEdge {
		opts = append(opts,
			chromedp.ExecPath("C:\\Program Files (x86)\\Microsoft\\Edge\\Application\\msedge.exe"),
		)
		Debug("Browser: Microsoft Edge")
	} else {
		Debug("Browser: Chrome (default)")
	}
	if definition.Secret {
		opts = append(opts,
			chromedp.Flag("incognito", true),
		)
	}
	if !definition.CertValidation {
		opts = append(opts,
			chromedp.Flag("ignore-certificate-errors", true),
		)
	}

	// Initialize contexts
	allocCtx, _ := chromedp.NewExecAllocator(context.Background(), opts...)
	runCtx, _ := chromedp.NewContext(allocCtx, chromedp.WithLogf(func(format string, v ...interface{}) {
		Debug("chromedp: "+format, v...)
	}))

	// Build actions based on the JSON line
	actions := []chromedp.Action{}
	for i, action := range definition.Actions {
		LogActionStart(i, action.Type, action.Target, action.Value)

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
			Info("Action %d: Sending account credentials", i)
		case "password":
			actions = append(actions,
				chromedp.SendKeys(action.Target,
					password,
					chromedp.ByQuery,
					chromedp.NodeVisible))
			Info("Action %d: Sending password credentials", i)
		case "sleep":
			actions = append(actions,
				chromedp.Sleep(time.Millisecond*time.Duration(action.Value)))
		default:
			Warn("Action %d: Unknown action type '%s'", i, action.Type)
		}

		LogActionComplete(i, action.Type)
	}

	Info("Executing %d chromedp actions", len(actions))
	err := chromedp.Run(runCtx, actions...)

	if err != nil {
		Error("Chromedp execution failed: %v", err)
		Fatal("Fatal error during chromedp execution")
		return false
	}

	LogChromedpComplete()
	return true
}
