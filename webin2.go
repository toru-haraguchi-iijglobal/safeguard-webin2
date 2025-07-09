// This program is a Website auto login program that work with
// RemoteApp-Launcher and One Identity Safeguard.
// It takes 4 parameters, those are:
//  1. JSON Lines (jsonl) filename which contains login actions
//     to each particular websites.
//     (-jsonl=<filename>)
//  2. Asset name which is the name of the website to login.
//     This parameter is passed by RemoteApp-Launcher.
//     (-asset=<name>)
//  3. Account ID to login to the website. This parameter
//     is passed by RemoteApp-Launcher.
//     (-account=<account>)
//  4. Password to login to the website. This parameter
//     is passed by RemoteApp-Launcher.
//     (-password=<password>)
package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

var jsonl_filename string
var yaml_filename string
var asset string
var account string
var password string

func init_args() bool {
	flag.StringVar(&jsonl_filename, "jsonl", "", "Specify the configuration JSON Lines file")
	flag.StringVar(&yaml_filename, "yaml", "", "Specify the configuration YAML file")
	flag.StringVar(&asset, "asset", "", "Specify the asset")
	flag.StringVar(&account, "account", "", "Specify the user")
	flag.StringVar(&password, "password", "", "Specify the password")
	flag.Parse()

	log.Printf("args: -jsonl=%s -yaml=%s -asset=%s -account=%s -password=%s",
		jsonl_filename, yaml_filename, asset, account, password)

	return true
}

func main() {
	// Open error logfile and redirect log
	exe_path, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	exe_path = filepath.Dir(exe_path)
	log_filename := exe_path + "\\webin2_" + strconv.Itoa(os.Getpid()) + ".log"
	logfile, err := os.OpenFile(log_filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer logfile.Close()
	log.SetOutput(logfile)

	// Process runtime arguments
	init_args()

	if jsonl_filename != "" && yaml_filename != "" {
		log.Printf("You can not specify both -jsonl and -yaml at a same time.")
		return
	}

	var definition Definition
	if jsonl_filename != "" {
		// Search the Asset passed from RemoteApp-Launcher
		definition = search_jsonlines(asset, jsonl_filename)
	}
	if yaml_filename != "" {
		definition = search_yaml(asset, yaml_filename)
	}

	if definition.Asset == "" {
		log.Printf("Asset \"%s\" not found in \"%s%s\"", asset, jsonl_filename, yaml_filename)
		logfile.Close()
		return
	}

	// run chromedp
	if run(definition) {
		// if chromedp runs correctly, remove the log file.
		logfile.Close()
		os.Remove(log_filename)
	}

}
