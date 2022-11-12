/*
@todo
*/

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const VERSION = "0.2.0"

var OUTPUT string = ""
var SHOULD_LOG bool = true
var SHOULD_WRITE bool = false
var SHOULD_CHECK_CONTACTS bool = false
var FOLDERS string
var MY_FOLDERS []string


func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

func init() {

	flag.StringVar(&FOLDERS, "i", "", "Please specify the input path(s)")
	flag.StringVar(&OUTPUT, "o", "", "Please specify the output file")
	flag.Bool("c", false, "Count contacts")
	flag.Bool("w", true, "Write changes")

	flag.Parse()

	if FOLDERS == "" {
		fmt.Println("Please specify the input folder(s)")
		os.Exit(1)
	}

	if OUTPUT != "" {
		SHOULD_LOG = true
		// logger.Init(OUTPUT)
	}

	if isFlagPassed("w") {
		SHOULD_WRITE = true
	}

	if isFlagPassed("c") {
		SHOULD_CHECK_CONTACTS = true
	}

	flag.Usage = func() {
		fmt.Println("Usage:")
		fmt.Println("./EXE -i \"<path> [paths]... [-o log file.json] [-c (do contact counts)] [-w dry run off, doing changes to filesystem]")
	}
}

func main() {

	var actions [][]string

	for _, folder := range strings.Split(FOLDERS, " ") {

		files, err := ioutil.ReadDir(folder)

		if err != nil {
			log.Fatal(err)
		}

		actions = append(actions, CheckRundowns(folder, files))

	}
	for _, action := range actions {
		for _, item := range action {
			fmt.Println(item)
		}
	}
}
