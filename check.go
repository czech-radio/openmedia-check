// Parse the week and year number from the given rundown file.
// Return the zero values if they are not found.

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	//	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

// Data holds Date, Week and File strings
type Data struct {
	Date string `json:"date"`
	Week string `json:"week"`
	File string `json:"file"`
	Dest string `json:"dest"`
}

// Message holds Index, Status strings and Data struct.
type Message struct {
	Index  int    `json:"index"`
	Status string `json:"status"`
	Action string `json:"action"`
	Data   Data   `json:"data"`
	// TODO (optional) Errors?
}

//----------------------------------------------------------------------------
//  RUNDOWNS
//----------------------------------------------------------------------------

// ParseRundown parses openmedia file and returns date.
func ParseRundown(handle io.Reader) (Year, Month, Day, Week int) {

	var year, month, day, week = 0, 0, 0, 0
	var scanner bufio.Scanner

	scanner = *bufio.NewScanner(transform.NewReader(handle, unicode.UTF16(unicode.LittleEndian, unicode.UseBOM).NewDecoder()))

	for scanner.Scan() {
		var line = fmt.Sprintln(scanner.Text())

		if strings.Contains(line, `FieldID = "1004"`) {
			reg := regexp.MustCompile("([0-9][0-9][0-9][0-9]{1})([0-9]{2})([0-9]{2})(T)")
			res := reg.FindStringSubmatch(line)

			date, err := time.Parse("20060102", res[1]+res[2]+res[3])

			if err != nil {
				log.Fatal(err)
			}

			year, month, day = date.Year(), int(date.Month()), date.Day()
			year, week = date.ISOWeek()
			break // Find first ocurrance!
		}
	}

	return year, month, day, week
}

// ReportRundowns takes path and filelist and outputs Message report.
func ReportRundowns(annova string, path string, files []os.FileInfo) []Message {

	var result = make([]Message, len(files))

	status := (map[bool]string{true: "SUCCESS", false: "FAILURE"})
	actions := (map[int]string{0: "none", 1: "mv", 2: "rm"})
	var actionNo int

	for i, file := range files {

		fext := filepath.Ext(file.Name())

		// File should be skipped because it is a directory or has wrong filename.
		if file.IsDir() || fext != ".xml" {
			continue // should it be logged, or other action executed?
		}

		fptr, err := os.Open(filepath.Join(path, file.Name()))

		if err != nil {
			log.Fatal(err)
		}

		if !strings.Contains(file.Name(), "RD") {
			continue
		}

		year, month, day, fileWeek := ParseRundown(fptr)
		dirWeek, _ := strconv.Atoi(filepath.Base(path)[1:])
		isFilenameTheSame := file.Name() == FixFilename(file.Name())

		//defer fptr.Close()
		fptr.Close() // defer leaves too many files opened at once

		if fileWeek == dirWeek && isFilenameTheSame {
			actionNo = 0
		} else if fileWeek != dirWeek || !isFilenameTheSame {
			actionNo = 1
		}

		message := &Message{
			Index:  i,
			Status: (status[fileWeek == dirWeek && isFilenameTheSame]),
			Action: actions[actionNo],
			Data: Data{
				Date: fmt.Sprintf("%04d-%02d-%02d", year, month, day),
				Week: fmt.Sprintf("W%02d", fileWeek),
				File: fmt.Sprintf("%s", filepath.Join(path, file.Name())),
				Dest: fmt.Sprintf(filepath.Join(fmt.Sprintf("%s", annova), "Rundowns", fmt.Sprintf("%04d", year), fmt.Sprintf("W%02d", fileWeek), FixFilename(file.Name()))),
			},
		}

		result = append(result, *message)

		FormatMessage(*message)

	}

	return result
}

// FormatMessage formats Message to JSON.
func FormatMessage(report Message) string {
	reportJSONLine, err := json.Marshal(report)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(reportJSONLine)) // How to send this to another function (Python yield style)?
	return string(reportJSONLine)
}

//----------------------------------------------------------------------------
// CONTACTS (TODO)
//----------------------------------------------------------------------------

// ParseContact get io.Reader handler and do open media contact counts.
func ParseContact(filename string) (Year, Month, Day, Week int) {

	//scanner := bufio.NewScanner(transform.NewReader(handle, unicode.UTF16(unicode.LittleEndian, unicode.UseBOM).NewDecoder()))

	split := strings.Split(filename, "_")

	date, err := time.Parse("20060102", split[len(split)-1][0:8])

	if err != nil {
		log.Fatal(err)
	}
	var year, month, day, week = 0, 0, 0, 0

	year, month, day = date.Year(), int(date.Month()), date.Day()
	year, week = date.ISOWeek()

	return year, month, day, week
}

// ReportContacts detects right date from CT files.
func ReportContacts(annova string, path string, files []os.FileInfo) []Message {
	var result = make([]Message, len(files))

	status := (map[bool]string{true: "SUCCESS", false: "FAILURE"})
	actions := (map[int]string{0: "none", 1: "mv", 2: "rm"})
	var actionNo int

	for i, file := range files {

		fext := filepath.Ext(file.Name())

		// File should be skipped because it is a directory or has wrong filename.
		if file.IsDir() || fext != ".xml" {
			continue // should it be logged, or other action executed?
		}

		fptr, err := os.Open(filepath.Join(path, file.Name()))

		if err != nil {
			log.Fatal(err)
		}

		if !strings.Contains(file.Name(), "CT") {
			continue
		}

		year, month, day, fileWeek := ParseContact(file.Name())
		dirWeek, _ := strconv.Atoi(filepath.Base(path)[1:])
		isFilenameTheSame := file.Name() == FixFilename(file.Name())

		// don't use defer, close file after read
		fptr.Close()

		if fileWeek == dirWeek && isFilenameTheSame {
			actionNo = 0
		} else if fileWeek != dirWeek || !isFilenameTheSame {
			actionNo = 1
		}

		message := &Message{
			Index:  i,
			Status: (status[fileWeek == dirWeek && isFilenameTheSame]),
			Action: actions[actionNo],
			Data: Data{
				Date: fmt.Sprintf("%04d-%02d-%02d", year, month, day),
				Week: fmt.Sprintf("W%02d", fileWeek),
				File: fmt.Sprintf("%s", filepath.Join(path, file.Name())),
				Dest: fmt.Sprintf(filepath.Join(fmt.Sprintf("%s", annova), "Contacts", fmt.Sprintf("%04d", year), fmt.Sprintf("W%02d", fileWeek), FixFilename(file.Name()))),
			},
		}

		result = append(result, *message)

		FormatMessage(*message)

	}

	return result

}

// RepairFiles do filechanges to files on disk.
func RepairFiles(actions []Message, shouldWriteChanges bool) {
	// Execute the commands stored in actions.
	for _, action := range actions {
		if action.Action == "mv" && shouldWriteChanges {

			//check whatever dest directory exists, if not create it
			dir, _ := filepath.Split(action.Data.Dest)
			_, err := os.Stat(dir)
			if os.IsNotExist(err) {
				err := os.Mkdir(dir, 0775)
				if err != nil {
					log.Fatal(err)
				}
			}

			// move file to
			e := os.Rename(action.Data.File, action.Data.Dest)
			if e != nil {
				log.Fatal(e)
			}
		}
	}
}

// RemoveEmptyLines should remove empty lines from file list.
func RemoveEmptyLines(annova string, path string, files []os.FileInfo) {

	for _, file := range files {
		fext := filepath.Ext(file.Name())

		// File should be skipped because it is a directory or has wrong filename.
		if file.IsDir() || fext != ".xml" {
			continue // should it be logged, or other action executed?
		}

		fptr, _ := os.Open(filepath.Join(path, file.Name()))
		scanner := bufio.NewScanner(transform.NewReader(fptr, unicode.UTF16(unicode.LittleEndian, unicode.UseBOM).NewDecoder()))

		defer fptr.Close()

		var modded []string
		for scanner.Scan() {
			var line = fmt.Sprintln(scanner.Text())
			// strings.Contains(`Is Empty="yes"`) ? mozno zredukovat hodne prazdnych poli
			if strings.TrimSpace(line) == "" {
				continue
			} else {
				modded = append(modded, line)
			}
		}

		/* TODO Minification writing new file in UTF-16 as it was read. */

	}
}

// FixFilename fixes filenames to standard form.
func FixFilename(orig string) string {
	var modified string = orig

	switch {
	case strings.Contains(orig, "_Pond??l??_"):
		modified = strings.Replace(orig, "_Pond??l??_", "_Mon_", -1)
	case strings.Contains(orig, "_??ter??_"):
		modified = strings.Replace(orig, "_??ter??_", "_Tue_", -1)
	case strings.Contains(orig, "_St??eda_"):
		modified = strings.Replace(orig, "_St??eda_", "_Wed_", -1)
	case strings.Contains(orig, "_??tvrtek_"):
		modified = strings.Replace(orig, "_??tvrtek_", "_Thu_", -1)
	case strings.Contains(orig, "_P??tek_"):
		modified = strings.Replace(orig, "_P??tek_", "_Fri_", -1)
	case strings.Contains(orig, "_Sobota_"):
		modified = strings.Replace(orig, "_Sobota_", "_Sat_", -1)
	case strings.Contains(orig, "_Ned??le_"):
		modified = strings.Replace(orig, "_Ned??le_", "_Sun_", -1)
	case strings.Contains(orig, "_Po_"):
		modified = strings.Replace(orig, "_Po_", "_Mon_", -1)
	case strings.Contains(orig, "_??t_"):
		modified = strings.Replace(orig, "_??t_", "_Tue_", -1)
	case strings.Contains(orig, "_St_"):
		modified = strings.Replace(orig, "_St_", "_Wed_", -1)
	case strings.Contains(orig, "_??t_"):
		modified = strings.Replace(orig, "_??t_", "_Thu_", -1)
	case strings.Contains(orig, "_P??_"):
		modified = strings.Replace(orig, "_P??_", "_Fri_", -1)
	case strings.Contains(orig, "_So_"):
		modified = strings.Replace(orig, "_So_", "_Sat_", -1)
	case strings.Contains(orig, "_Ne_"):
		modified = strings.Replace(orig, "_Ne_", "_Sun_", -1)
	}

	modified = strings.Replace(modified, "__", "_", -1)

	return modified
}
