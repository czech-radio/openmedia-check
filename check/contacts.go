package check

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

func countContacts(filename string) (int, error) {
	// IO!
	file, err := os.Open(filename)
	if err != nil {
		// logger.Fatal("Error reading file: " + filename)
	}
	var count int = 0

	scanner := bufio.NewScanner(
		transform.NewReader(
			file, unicode.UTF16(unicode.LittleEndian, unicode.UseBOM).NewDecoder()))
	for scanner.Scan() {
		var line string = fmt.Sprintln(scanner.Text())
		if strings.Contains(line, "\"ContactContainerFieldID\" IsEmpty = \"no\"") {
			count += 1
		}
	}

	if err := file.Close(); err != nil {
		// logger.Fatal(err.Error())
		return -1, err
	}

	if count == -1 {
		// logger.Fatal("Error processing data")
		return -1, err
	}

	return count, err
}

func CheckContacts(folder string) error {

	total := 0

	var errornous_filenames []string

	files, err := ioutil.ReadDir(folder)
	if err != nil {
		// logger.Fatal(err.Error())
		return err
	}

	for _, fn := range files {
		if strings.Contains(fn.Name(), ".xml") {
			contacts, err := countContacts(filepath.Join(folder, fn.Name()))
			if err != nil {
				log.Fatal(err.Error())
			}
			total += contacts
			total++
		} else {
			errornous_filenames = append(errornous_filenames, "Not a xml file: "+fn.Name())
		}
	}
	// logger.Println("No. of contacts collected: " + fmt.Sprint(contactsTotal) + " SUCCESS!")
	return nil
}
