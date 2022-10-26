package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"github.com/beevik/etree"
)

const VERSION = "0.0.2"

var ROOT_DIR string = ""
var OUTPUT string = ""
var LOG bool = false
var logfile os.File

func init() {
	flag.StringVar(&ROOT_DIR, "i", "", "Please specify the input path")
	flag.StringVar(&OUTPUT, "o", "", "Please specify the output file")
	flag.Parse()

	if ROOT_DIR == "" {
		log.Fatal("Please specify the input folder -i ../..")
		return
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
		fmt.Println("./openmedia-files-checker -i /path/to/openmedia/Rundown (full path to Rundowns folder) [-o logfile.txt]")
	}
}

/*
type OM_Object struct {
	XMLName xml.Name   `xml:"OM_OBJECT"`
	Fields  []OM_Field `xml:"OM_FIELD"`
}

type OM_Field struct {
	FieldID   string `xml:"FiledID,attr"`
	FieldType   string `xml:"FiledType,attr"`
	FieldName   string `xml:"FiledName,attr"`
}
*/

func get_contact_count(filename string) int {
        log.Println("Processing file "+filename)
        doc := etree.NewDocument()
        if err:= doc.ReadFromFile(filename); err != nil{
		log.Fatal("Error reading file " + filename)
	}
        
        var count int = 0

        for _, e := range doc.FindElements("./OM_OBJECT/*"){
          log.Println(e.Tag)
          count += 1
        }
        
        log.Println("ok got " + fmt.Sprint(count))

	return count
}

func main() {
	years, err := ioutil.ReadDir(ROOT_DIR)
	if err != nil {
		log.Fatal("Input folder cannot be opened.")
	}
	// years list
	for _, year := range years {

		log.Println("Checking Rundown count Year " + year.Name())

		files, err := ioutil.ReadDir(ROOT_DIR + "/" + year.Name())
		if err != nil {
			log.Fatal(err)
		}

		// week list
		for _, f := range files {

			checked := 0
			var errornous_filenames []string
                        var contacts int = 0

			files, err := ioutil.ReadDir(ROOT_DIR + "/" + year.Name() + "/" + f.Name())
			if err != nil {
				log.Fatal(err)
			} // files list
			for _, fn := range files {
				if strings.Contains(fn.Name(), "xml") {

					week_no, err := strconv.Atoi(strings.Split(fmt.Sprint(fn.ModTime().ISOWeek()), " ")[1])
					if err != nil {
						log.Fatal(err)
					}

					dir_no, err := strconv.Atoi(strings.Split(fmt.Sprint(f.Name()), "W")[1])
					if err != nil {
						log.Fatal(err)
					}

					if week_no == dir_no {
						checked += 1
					} else {
						errornous_filenames = append(errornous_filenames, "Wrong modtime descriptor in file: "+f.Name()+"/"+fn.Name())
						// log.Printf("file was modded on: %v and is in dir %v (%s)",week_no,dir_no,fn.Name())
					}

					contacts += get_contact_count(ROOT_DIR + "/" + year.Name() + "/" + f.Name() + "/" + fn.Name())

				} else {
					errornous_filenames = append(errornous_filenames, "Not a xml file: "+f.Name()+"/"+fn.Name())
				}
			}

			if checked == len(files) {
				log.Println(year.Name() + "/" + f.Name() + ": comparing file modtime to foldername: " + fmt.Sprint(checked) + "/" + fmt.Sprint(len(files)) + " PASSED!")

			} else {
				log.Println(year.Name() + "/" + f.Name() + ": comparing file modtime to foldername: " + fmt.Sprint(checked) + "/" + fmt.Sprint(len(files)) + " NOT PASSED!")

				for _, ef := range errornous_filenames {
					log.Println("mismatch found: " + fmt.Sprint(ef))
				}
			}
		}
	}

	if LOG {
		defer logfile.Close()
	}

}
