// Parse the week and year number from the given rundown file.
// Return the zero values if they are not found.

package check

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

func ParseRundownDate(handle io.Reader) (int, int, int, int) {

	var year, month, day, week = 0, 0, 0, 0

	scanner := bufio.NewScanner(transform.NewReader(handle, unicode.UTF16(unicode.LittleEndian, unicode.UseBOM).NewDecoder()))

	for scanner.Scan() {
		var line = fmt.Sprintln(scanner.Text())

		if !strings.Contains(line, "FieldID = \"1004\"") {
			continue
		}

		reg := regexp.MustCompile("(202[0-9]{1})([0-9]{2})([0-9]{2})(T)")
		res := reg.FindStringSubmatch(line)

		date, err := time.Parse("20060102", res[1]+res[2]+res[3])

		if err != nil {
			log.Fatal(err)
		}

		year, month, day = date.Year(), int(date.Month()), date.Day()
		_, week = date.ISOWeek()

		break
	}

	return year, month, day, week
}

func CheckRundowns(path string, files []os.FileInfo) []string {

	var actions []string

	for _, file := range files {

		fext := filepath.Ext(file.Name())

		if fext != ".xml" {
			// File shoud be removed.
			actions = append(actions, "REMOVE")
			continue
		}

		fptr, err := os.Open(filepath.Join(path, file.Name()))

		if err != nil {
			log.Fatal(err)
		}

		year, month, day, fileWeek := ParseRundownDate(fptr)

		dirWeek, _ := strconv.Atoi(filepath.Base(path)[1:])

		status := (map[bool]string{true: "SUCCESS", false: "FAILURE"})[fileWeek == dirWeek]

		fmt.Println(status, year, month, day, fileWeek, path, file.Name())
	}

	return actions
}

func PlaceRundows(actions []string) {
	// Execute the commands stored in actions
}
