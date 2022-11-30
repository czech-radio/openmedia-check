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
	"path"
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
func ParseRundown(handle io.Reader, isUtf16 bool) (Year, Month, Day, Week int) {

	var year, month, day, week = 0, 0, 0, 0
	var scanner bufio.Scanner

	if isUtf16 {
		scanner = *bufio.NewScanner(transform.NewReader(handle, unicode.UTF16(unicode.LittleEndian, unicode.UseBOM).NewDecoder()))
	} else {
		scanner = *bufio.NewScanner(transform.NewReader(handle, unicode.UTF8.NewDecoder()))
	}

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

		defer fptr.Close()

		year, month, day, fileWeek := ParseRundown(fptr, true)
		dirWeek, _ := strconv.Atoi(filepath.Base(path)[1:])

		if fileWeek == dirWeek {
			actionNo = 0
		} else if fileWeek != dirWeek {
			actionNo = 1
		}

		message := &Message{
			Index:  i,
			Status: (status[fileWeek == dirWeek]),
			Action: actions[actionNo],
			Data: Data{
				Date: fmt.Sprintf("%04d-%02d-%02d", year, month, day),
				Week: fmt.Sprintf("W%02d", fileWeek),
				File: fmt.Sprintf("%s", filepath.Join(path, file.Name())),
				Dest: fmt.Sprintf(filepath.Join(fmt.Sprintf("%s", annova), "Rundowns", fmt.Sprintf("%04d", year), fmt.Sprintf("W%02d", fileWeek))),
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

// RepairRundows (unimplemented) do filechanges to files on disk.
func RepairRundows(actions []Message, shouldWriteChanges bool) {
	// Execute the commands stored in actions.
	for _, action := range actions {
		if action.Action == "mv" && shouldWriteChanges {

			//check whatever dest directory exists, if not create it
			_, err := os.Stat(action.Data.Dest)
			if os.IsNotExist(err) {
				err := os.Mkdir(action.Data.Dest, 0775)
				if err != nil {
					log.Fatal(err)
				}
			}

			// move file to
			_, filename := filepath.Split(action.Data.File)
			e := os.Rename(action.Data.File, path.Join(action.Data.Dest, filename))
			if e != nil {
				log.Fatal(e)
			}
		}
	}
}

//----------------------------------------------------------------------------
// CONTACTS (TODO)
//----------------------------------------------------------------------------

// ParseContact get io.Reader handler and do open media contact counts.
func ParseContact(handle io.Reader) int {

	scanner := bufio.NewScanner(transform.NewReader(handle, unicode.UTF16(unicode.LittleEndian, unicode.UseBOM).NewDecoder()))

	var count int = 0
	for scanner.Scan() {
		// TODO `"ContactContainerFieldID IsEmpty = "no"`
	}

	return count
}

// ReportContacts makes openmedia contacts count and outputs slice of ... (unimplemented)
func ReportContacts(path string, files []os.FileInfo) []string {
	var result = make([]string, len(files))

	/* TODO */

	return result

}

// RepairContacts fixes ... (unimplemented)
func RepairContacts(actions []string) {
	// TODO Execute the commands stored in actions.
}
