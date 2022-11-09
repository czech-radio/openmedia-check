package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"path/filepath"

	"io/ioutil"
	"os"

	"strconv"
	"strings"
	"time"

	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

//// SCOPE ////////////////////////////////////////////////////////

const VERSION = "0.0.2"

var logger JSONLogger

var OUTPUT string = ""
var LOG bool = false
var DRY_RUN bool = true
var CONTACTS bool = false
var FOLDERS string
var MY_FOLDERS []string

type JSONformat struct {
	Date string ``
}

func init() {
	logger = JSONLogger{Filename: "log.json"}
	parseArguments()
}

func main() {
	for _, FOLDER := range MY_FOLDERS {

		// main check here
		err := check_files_inner_date_to_foldername(FOLDER)
		if err != nil {
			logger.Fatal(err.Error())
		}

		// optional checking contact count
		if CONTACTS {
			err := check_contact_count(FOLDER)
			if err != nil {
				logger.Fatal(err.Error())
			}
		}

	} // end range FOLDERS

	if LOG {
		defer logger.Write()
	}
}

func parseArguments() {
	flag.StringVar(&FOLDERS, "i", "", "Please specify the input path(s)")
	flag.StringVar(&OUTPUT, "o", "", "Please specify the output file")
	flag.Bool("c", false, "Count contacts")
	flag.Bool("w", true, "Write changes")
	//flag.CommandLine.SetOutput(io.Discard)
	flag.Parse()

	if FOLDERS == "" {
		logger.Fatal("Please specify the input folder(s) -i /path/to/2022/W33")
	} else {
		
               MY_FOLDERS = strings.Split(FOLDERS, " ")
        }

	if OUTPUT != "" {
		LOG = true
		logger.Init(OUTPUT)
	}

	if isFlagPassed("c") {
		CONTACTS = true
	}

	if isFlagPassed("w") {
		DRY_RUN = false
	}

	flag.Usage = func() {
		fmt.Println("Usage of program:")
		fmt.Println("./openmedia_files_checker -i \"/path/to/Rundown1 /path/to/Rundown2\" (full path(s) to Rundowns folder(s)) [-o log file.json] [-c (do contact counts)] [-w dry run off, doing changes to filesystem]")
	}
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

func removeEmptyStrings(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func parseWeekNumber(filename string) (int, int, error) {
	file, err := os.Open(filename)
	if err != nil {
		logger.Fatal("Error reading file: " + filename)
	}

	var Year int = -1
	var week int = -1

	scanner := bufio.NewScanner(transform.NewReader(file, unicode.UTF16(unicode.LittleEndian, unicode.UseBOM).NewDecoder()))
	for scanner.Scan() {
		var line string = fmt.Sprintln(scanner.Text())
		var offset = strings.Index(line, "<OM_DATETIME>")
		if offset != -1 && strings.Contains(line, "\"Čas začátku\" IsEmpty = \"no\"") {

			offset2 := 13 // len of <OM_DATETIME> string
			dateInt, _ := strconv.Atoi(line[offset+offset2+6 : offset+offset2+8])
			month, _ := strconv.Atoi(line[offset+offset2+4 : offset+offset2+6])
			year, _ := strconv.Atoi(line[offset+offset2 : offset+offset2+4])
			then := time.Date(year, time.Month(month), dateInt, 0, 0, 0, 0, time.UTC)
			Year, week = then.ISOWeek()

			// get first only?
			break
		}
	}
	err = file.Close()
	if err != nil {
		logger.Fatal(err.Error())
	}
	return Year, week, err
}

func countContacts(filename string) (int, error) {

	file, err := os.Open(filename)
	if err != nil {
		logger.Fatal("Error reading file: " + filename)
	}
	var count int = 0

	scanner := bufio.NewScanner(transform.NewReader(file, unicode.UTF16(unicode.LittleEndian, unicode.UseBOM).NewDecoder()))
	for scanner.Scan() {
		var line string = fmt.Sprintln(scanner.Text())
		if strings.Contains(line, "\"ContactContainerFieldID\" IsEmpty = \"no\"") {
			count += 1
		}
	}

	if err := file.Close(); err != nil {
		logger.Fatal(err.Error())
		return -1, err
	}

	if count == -1 {
		logger.Fatal("Error processing data")
		return -1, err
	}

	return count, err
}

func folder_name_to_new_one(folder string, year int, month int) string {
	split := strings.Split(folder, "/")

	split[len(split)-1] = fmt.Sprintf("W%02d", month)
	split[len(split)-2] = fmt.Sprintf("%04d", year)

	var newpath string
	if folder[0:1] == "/" {
		newpath = "/"
	} else {
		newpath = ""
	}

	for _, i := range split {
		newpath = filepath.Join(newpath, i)
	}

	return newpath
}

func check_files_inner_date_to_foldername(FOLDER string) error {

	//foldername := filepath.Base(FOLDER)
	files, err := ioutil.ReadDir(FOLDER)
	if err != nil {
		logger.Fatal(err.Error())
		return err
	}

	var errornous_filenames []string
	var count int

	for _, fn := range files {
		if strings.Contains(fn.Name(), ".xml") {
			year, week_no, err := parseWeekNumber(filepath.Join(FOLDER, fn.Name()))

			if err != nil {
				logger.Fatal(err.Error())
			}

			dir_no, err := strconv.Atoi(strings.Split(fmt.Sprint(FOLDER), "W")[1])
			if err != nil {
				logger.Fatal(err.Error())
			}

			if week_no == dir_no {
				count += 1
			} else {
				errornous_filenames = append(errornous_filenames, "mv "+filepath.Join(FOLDER, fn.Name())+" "+folder_name_to_new_one(FOLDER, year, week_no)+"/"+fn.Name())
			}
		} else if !fn.IsDir() {
			errornous_filenames = append(errornous_filenames, "rm "+filepath.Join(FOLDER, fn.Name()))

		}
	}

	if count == len(files) {
		logger.Println(FOLDER + " test result: " + fmt.Sprint(count) + "/" + fmt.Sprint(len(files)) + " SUCCESS!")
	} else {
		logger.Fatal(FOLDER + " test result: " + fmt.Sprint(count) + "/" + fmt.Sprint(len(files)) + " FAILURE!")
		for _, ef := range errornous_filenames {
			if DRY_RUN {
				logger.Warn("DRY_RUN ON (not) doing: " + fmt.Sprint(ef))
			} else {
				command := strings.Split(fmt.Sprint(ef), " ")

				logger.Warn("SHARP_MODE ON doing: " + fmt.Sprint(ef))

				// move files
				if command[0] == "mv" {

					err := os.Rename(command[1], command[2])
					if err != nil {
						logger.Fatal(err.Error())
					}
					//remove files
				} else if command[0] == "rm" {
					err := os.Remove(command[1])
					if err != nil {
						logger.Fatal(err.Error())
					}
				}
			}
		}
	}

	return nil
}

func check_contact_count(FOLDER string) error {

	contactsTotal := 0
	checked := 0

	var errornous_filenames []string

	files, err := ioutil.ReadDir(FOLDER)
	if err != nil {
		logger.Fatal(err.Error())
		return err
	}

	for _, fn := range files {
		if strings.Contains(fn.Name(), ".xml") {
			contacts, err := countContacts(filepath.Join(FOLDER, fn.Name()))
			if err != nil {
				logger.Fatal(err.Error())
			}
			contactsTotal += contacts
			checked++
		} else {
			errornous_filenames = append(errornous_filenames, "Not a xml file: "+fn.Name())
		}
	}
	logger.Println("No. of contacts collected: " + fmt.Sprint(contactsTotal) + " SUCCESS!")
	return nil
}
