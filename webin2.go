// This program is a Website auto login program that work with
// RemoteApp-Launcher and One Identity Safeguard.
// It takes 4 parameters, those are:
//  1. JSON Lines (jsonl) filename which contains login actions
//     to each particular websites.
//     (-jsonl=<filename>)
//  2. Asset name which is the name of the website to login.
//     This parameter is passed by RemoteApp-Launcher.
//     (-asset=<name>)
//  3. User ID to login to the website. This parameter
//     is passed by RemoteApp-Launcher.
//     (-user=<username>)
//  4. Password to login to the website. This parameter
//     is passed by RemoteApp-Launcher.
//     (-pwd=<password>)
package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/chromedp/chromedp"
)

var jsonl_filename string
var asset string
var user string
var password string

func init_args() bool {
	flag.StringVar(&jsonl_filename, "jsonl", "", "Specify the configuration JSON Lines file")
	flag.StringVar(&asset, "asset", "", "Specify the asset")
	flag.StringVar(&user, "user", "", "Specify the user")
	flag.StringVar(&password, "pwd", "", "Specify the password")

	return true
}

func main() {
	// Open error logfile and redirect log
	exe_path, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	exe_path = filepath.Dir(exe_path)
	log_filename := exe_path + "\\webin_" + strconv.Itoa(os.Getpid()) + ".log"
	logfile, err := os.OpenFile(log_filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer logfile.Close()
	log.SetOutput(logfile)
	// Process runtime arguments
	init_args()
	// Open JSONL file
	// TODO: Load the JSONL file and parse it into a variable of type JsonLine
	var json_line JsonLine = search_jsonlines(jsonl_filename)

	// Search the Asset passed from RemoteApp-Launcher
	// If match
	// prepare a chromedp action array
	var actions []chromedp.Action = build_actions(json_line)
	if len(actions) > 0 {
		log.Printf("%d actions found for asset: %s", len(actions), asset)
	} else {
		// else log error and exit.
		log.Printf("%d actions found for asset: %s", len(actions), asset)
	}
	// run chromedp
	logfile.Close()
	// if chromedp runs correctly, remove the log file.
	if run(actions) {
		os.Remove(log_filename)
	}
}
