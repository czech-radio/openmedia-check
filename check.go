// Parse the week and year number from the given rundown file.
// Return the zero values if they are not found.

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

//----------------------------------------------------------------------------
//  RUNDOWNS
//----------------------------------------------------------------------------

func ParseRundown(handle io.Reader) (int, int, int, int) {

	var year, month, day, week = 0, 0, 0, 0

	scanner := bufio.NewScanner(transform.NewReader(handle, unicode.UTF16(unicode.LittleEndian, unicode.UseBOM).NewDecoder()))

	for scanner.Scan() {
		var line = fmt.Sprintln(scanner.Text())

		if strings.Contains(line, `FieldID = "1004"`) {
			reg := regexp.MustCompile("(202[0-9]{1})([0-9]{2})([0-9]{2})(T)")
			res := reg.FindStringSubmatch(line)

			date, err := time.Parse("20060102", res[1]+res[2]+res[3])

			if err != nil {
				log.Fatal(err)
			}

			year, month, day = date.Year(), int(date.Month()), date.Day()
			_, week = date.ISOWeek()
			break; // Find first occurence!
		}
	}

	return year, month, day, week
}

func CheckRundowns(path string, files []os.FileInfo) []string {

	var result = make([]string, len(files))

	status := (map[bool]string{true: "SUCCESS", false: "FAILURE"})

	sem := make(chan struct{}, len(files))

	for i, file := range files {

		// File shoud be moved because it is a directory.
		if file.IsDir() {
			result = append(result, `{"status": "FAILURE", "action": "MOVE"`+", file: "+file.Name()+"}")
			log.Println("SKIPPING")
			continue
		}

		fext := filepath.Ext(file.Name())

		// File shoud be moved because is has wrong file extension.
		if fext != ".xml" {
			result = append(result, `{"status": "FAILURE", "action": "MOVE"`+", file: "+file.Name()+"}")
			log.Println("SKIPPING")
			continue
		}

		fptr, err := os.Open(filepath.Join(path, file.Name()))

		if err != nil {
			log.Fatal(err)
			log.Println("ERROR")
		}

		// go func(i int) { }(i)
		year, month, day, fileWeek := ParseRundown(fptr)
		dirWeek, _ := strconv.Atoi(filepath.Base(path)[1:])
		result = append(result, `{"status": "`+status[fileWeek == dirWeek]+`", "data": {"date": "`+fmt.Sprint(year)+"-"+fmt.Sprint(month)+"-"+fmt.Sprint(day)+`", "week": `+fmt.Sprint(fileWeek)+`, "file": "`+file.Name()+`"} }`)
		sem <- struct{}{}
		defer fptr.Close()

	}

	// sem.Wait(len(files)); // FIXME

	return result
}

func PlaceRundows(actions []string) {
	// Execute the commands stored in actions.
}

//----------------------------------------------------------------------------
// CONTACTS (TODO)
//----------------------------------------------------------------------------

func parseContact(handle io.Reader) {

	scanner := bufio.NewScanner(transform.NewReader(handle, unicode.UTF16(unicode.LittleEndian, unicode.UseBOM).NewDecoder()))

	for scanner.Scan() {
		// `"ContactContainerFieldID IsEmpty = "no"`
	}

	return
}

func CheckContacts(path string, files []os.FileInfo) []string {
	var result = make([]string, len(files))
	return result

}
