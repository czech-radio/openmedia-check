/*
@todo
*/

package main

import (
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// VERSION of a program
const VERSION = "0.2.0"

// ShouldWriteChanges is a switch to write changes on disk or not true : false.
var ShouldWriteChanges bool

// ShouldCheckContacts is a switch to check contacts in file or not, true : false.
var ShouldCheckContacts bool

// ANNOVA is SYSVAR pointing to Openmedia root folder path
var ANNOVA string

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

	ANNOVA := os.Getenv("ANNOVA")

	if ANNOVA == "" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("System has no variable $ANNOVA neither proper .env file is present.")
		}
	}

	INPUTS := flag.String("i", "", "The input directories.")
	OUTPUT := flag.String("o", "", "The output file name.")

	ShouldWriteChanges := flag.Bool("w", false, "Should write changes?")
	ShouldCheckContacts := flag.Bool("c", false, "Should count contacts?")

	flag.Parse()

	if *INPUTS == "" {
		fmt.Println("Please specify the input folder(s)")
		os.Exit(1)
	}

	if *OUTPUT != "" {
		// Write to stdout
	}

	if isFlagPassed("w") {
		*ShouldWriteChanges = true
	}

	if isFlagPassed("c") {
		*ShouldCheckContacts = true
	}

	flag.Usage = func() {
		fmt.Println("Usage:")
		fmt.Println(`./openmedia-check -i "<path> [path...]" [-o <file_name>] [-c] [-w]`)
	}

	var actions [][]Message

	for _, folder := range strings.Split(*INPUTS, " ") {

		files, err := ioutil.ReadDir(folder)

		if err != nil {
			log.Fatal(err)
		}

		actions = append(actions, ReportRundowns(ANNOVA, folder, files))
	}
}
