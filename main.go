package main

import (
	"bufio"
	"errors"
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
	"time"

	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

//// TODO  ////////////////////////////////////////////////////////
// detect naming function
// map of suggested moves
// json logging
//////////////////////////

//// SCOPE ////////////////////////////////////////////////////////

const VERSION = "0.0.2"

var CONTACTS bool

// var FOLDER string = ""
var OUTPUT string = ""
var LOG bool = false
var logfile os.File

var FOLDERS string
var MY_FOLDERS []string

type JSONformat struct {
	Date string ``
}

//// INIT /////////////////////////////////////////////////////////

func init() {
	parse_args()
}

//// MAIN /////////////////////////////////////////////////////////

func main() {
	for _, FOLDER := range MY_FOLDERS {

		err := check_files_filename_to_foldername(FOLDER)
		if err != nil {
			log.Fatal(err)
		}

		err = check_files_moddtime_to_foldername(FOLDER)
		if err != nil {
			log.Fatal(err)
		}
	} // end range FOLDERS

	if LOG {
		defer logfile.Close()
	}
}

//// FUNCTIONS ////////////////////////////////////////////////////

func parse_args() {

	flag.StringVar(&FOLDERS, "i", "", "Please specify the input path(s)")
	flag.StringVar(&OUTPUT, "o", "", "Please specify the output file")
	flag.Parse()

	if FOLDERS == "" {
		log.Fatal("Please specify the input folder(s) -i /path/to/2022/W33")
	} else {
		MY_FOLDERS = strings.Split(FOLDERS, " ")
		log.Println(FOLDERS)
	}

	if OUTPUT != "" {
		LOG = true
		log.Println("Appending to logfile: " + OUTPUT)
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

func filename_to_weekno(filename string) (int, error) {

	if filename == "" {
		log.Fatal("No filename was supplied")
		return 0, errors.New("The filename do not exist")
	} else {

		split1 := strings.Split(filename, "-")
		ss := strings.Split(split1[len(split1)-1], "_")

		end := ss[len(ss)-1]

		dateInt, _ := strconv.Atoi(end[6:8])
		month, _ := strconv.Atoi(end[4:6])
		year, _ := strconv.Atoi(end[0:4])
		then := time.Date(year, time.Month(month), dateInt, 0, 0, 0, 0, time.UTC)
		_, week := then.ISOWeek()

		return week, nil

	}

}

func get_contact_count(filename string) (int, error) {

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
		return -1, err
	}

	if count == -1 {
		log.Fatal("Error processing data")
		return -1, err
	}

	return count, err
}

func check_files_filename_to_foldername(FOLDER string) error {

	foldername := filepath.Base(FOLDER)
	files, err := ioutil.ReadDir(FOLDER)
	if err != nil {
		log.Fatal(err)
		return err
	}

	var errornous_filenames []string
	var count int

	for _, fn := range files {
		week_no, err := filename_to_weekno(fn.Name())
		if err != nil {
			log.Fatal(err)
		}

		dir_no, err := strconv.Atoi(strings.Split(fmt.Sprint(FOLDER), "W")[1])
		if err != nil {
			log.Fatal(err)
		}

		if week_no == dir_no {
			count += 1

		} else {
			log.Fatal(fn.Name() + " Is in wrong folder")
			errornous_filenames = append(errornous_filenames, "Wrong file placement: "+FOLDER+"/"+fn.Name())

		}
	}

	if count == len(files) {
		log.Println(foldername + ": Comparing filename dates to foldername: " + fmt.Sprint(count) + "/" + fmt.Sprint(len(files)) + " PASSED!")
	} else {
		log.Println(foldername + ": Comparing filename dates to foldername: " + fmt.Sprint(count) + "/" + fmt.Sprint(len(files)) + " NOT PASSED!")
		for _, ef := range errornous_filenames {
			log.Println("mismatch found: " + fmt.Sprint(ef))
		}
	}

	return nil
}

func check_files_moddtime_to_foldername(FOLDER string) error {

	checked := 0
	var errornous_filenames []string
	var contactsTotal int = 0

	foldername := filepath.Base(FOLDER)
	files, err := ioutil.ReadDir(FOLDER)
	if err != nil {
		log.Fatal(err)
		return errors.New("Cannot read directory")
	}

	//log.Println(foldername)

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
				contacts, err := get_contact_count(filepath.Join(FOLDER, fn.Name()))
				if err != nil {
					log.Fatal(err)
				}
				contactsTotal += contacts
				log.Println("No. of contacts collected: " + fmt.Sprint(contacts) + "/" + fmt.Sprint(contactsTotal))
			}

		} else {
			errornous_filenames = append(errornous_filenames, "Not a xml file: "+fn.Name())
		}

	} // end range files
	if checked == len(files) {
		log.Println(foldername + ": Comparing file modtime to foldername: " + fmt.Sprint(checked) + "/" + fmt.Sprint(len(files)) + "   PASSED!")
	} else {
		log.Println(foldername + ": Comparing file modtime to foldername: " + fmt.Sprint(checked) + "/" + fmt.Sprint(len(files)) + "   NOT PASSED!")
		for _, ef := range errornous_filenames {
			log.Println("mismatch found: " + fmt.Sprint(ef))
		}
	}

	return nil
}
