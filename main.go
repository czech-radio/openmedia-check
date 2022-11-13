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

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

func main() {

	INPUTS := flag.String("i", "", "The input directories.")
	OUTPUT := flag.String("o", "", "The output file name.")

	SHOULD_WRITE_CHANGES := flag.Bool("w", false, "Should write changes?")
	SHOULD_CHECK_CONTACTS := flag.Bool("c", false, "Should count contacts?")

	flag.Parse()

	if *INPUTS == "" {
		fmt.Println("Please specify the input folder(s)")
		os.Exit(1)
	}


	if *OUTPUT != "" {
		// Write to stdout
	}

	if isFlagPassed("w") {
		*SHOULD_WRITE_CHANGES = true
	}

	if isFlagPassed("c") {
		*SHOULD_CHECK_CONTACTS = true
	}

	flag.Usage = func() {
		fmt.Println("Usage:")
		fmt.Println(`./openmedia-check -i "<path> [path...]" [-o <file_name>] [-c] [-w]`)
	}

	var actions [][]string

	for _, folder := range strings.Split(*INPUTS, " ") {

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
