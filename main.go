package main

import (
	"bufio"
	"flag"
	"fmt"
	"path/filepath"

	//	"github.com/beevik/etree"
	//	"go.uber.org/zap"
	"io/ioutil"
	"log"

	"os"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

const VERSION = "0.0.2"

var CONTACTS bool

//// TODO ////////////////
// detect naming function
// map of suggested moves
// json logging
//////////////////////////

// var FOLDER string = ""
var OUTPUT string = ""
var LOG bool = false
var logfile os.File

var FOLDERS string
var MY_FOLDERS []string

type JSONformat struct {
	Date string ``
}

func init() {
	flag.StringVar(&FOLDERS, "i", "", "Please specify the input path(s)")
	flag.StringVar(&OUTPUT, "o", "", "Please specify the output file")
	flag.Parse()

	/*
	  logger, err := zap.NewProduction()
	  if err != nil {
	    log.Fatal(err)
	  }

	  //log := zerolog.New(os.Stdout).With().Timestamp.Logger()
	*/

	log.Println(FOLDERS)

	if FOLDERS == "" {
		log.Fatal("Please specify the input folder(s) -i /path/to/2022/W33")
		return
	} else {
		MY_FOLDERS = strings.Split(FOLDERS, " ")
	}

	if OUTPUT != "" {
		LOG = true
		log.Println("Creating logfile: " + OUTPUT)
		logfile, err := os.OpenFile(OUTPUT, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			log.Fatal("Cannot open output file for writing.")
		}
		log.SetOutput(logfile)
	}

	flag.Usage = func() {
		fmt.Println("Usage of program:")
		fmt.Println("./openmedia-files-checker -i /path/to/openmedia/Rundown (full path(s) to Rundowns folder(s)) [-o logfile.txt]")
	}
}

func filename_to_weekno(filename string, foldername string) bool {
	parsed := strings.Split(filename, "_")
	log.Println(parsed)
	return true
}

func get_contact_count(filename string) int {

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Error reading file: " + filename)
	}
	var count int = 0

	scanner := bufio.NewScanner(transform.NewReader(file, unicode.UTF16(unicode.LittleEndian, unicode.UseBOM).NewDecoder()))
	for scanner.Scan() {
		var line string = fmt.Sprintln(scanner.Text())
		if strings.Contains(line, "\"ContactContainerFieldID\" IsEmpty = \"no\"") {
			//fmt.Println(fmt.Sprint(count) + " " + line)
			count += 1
		}
	}

	if err := file.Close(); err != nil {
		log.Fatal(err)
	}

	if count == -1 {
		log.Fatal("Error processing data")
		return 0
	}

	return count
}

func main() {

	for _, FOLDER := range MY_FOLDERS {

		checked := 0
		var errornous_filenames []string
		var contacts int = 0
		var contactsTotal int = 0

		foldername := filepath.Base(FOLDER)
		files, err := ioutil.ReadDir(FOLDER)

		log.Println("Processing folder: " + FOLDER)
		//log.Println(foldername)

		if err != nil {
			log.Fatal(err)
		} // files list

		for _, fn := range files {
			if strings.Contains(fn.Name(), ".xml") {

				week_no, err := strconv.Atoi(strings.Split(fmt.Sprint(fn.ModTime().ISOWeek()), " ")[1])
				if err != nil {
					log.Fatal(err)
				}

				dir_no, err := strconv.Atoi(strings.Split(fmt.Sprint(foldername), "W")[1])
				if err != nil {
					log.Fatal(err)
				}

				if week_no == dir_no {
					checked += 1
				} else {
					errornous_filenames = append(errornous_filenames, "Wrong modtime descriptor in file: "+foldername+"/"+fn.Name())
					// log.Printf("file was modded on: %v and is in dir %v (%s)",week_no,dir_no,fn.Name())
				}

				if CONTACTS == true {
					contacts = get_contact_count(filepath.Join(FOLDER, fn.Name()))
					contactsTotal += contacts
					log.Println("No. of contacts collected: " + fmt.Sprint(contacts) + "/" + fmt.Sprint(contactsTotal))
				}

			} else {
				errornous_filenames = append(errornous_filenames, "Not a xml file: "+fn.Name())
			}

		} // end range files
		if checked == len(files) {
			log.Println(foldername + ": comparing file modtime to foldername: " + fmt.Sprint(checked) + "/" + fmt.Sprint(len(files)) + " PASSED!")
		} else {
			log.Println(foldername + ": comparing file modtime to foldername: " + fmt.Sprint(checked) + "/" + fmt.Sprint(len(files)) + " NOT PASSED!")
			for _, ef := range errornous_filenames {
				log.Println("mismatch found: " + fmt.Sprint(ef))
			}
		}
	} // end range FOLDERS
	if LOG {
		defer logfile.Close()
	}
} // end main
