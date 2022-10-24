package main

import (
        "fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
        "flag"
)

var years []string = []string{"2020","2021","2022"}
var ROOT_DIR string = ""

func init(){
  flag.StringVar(&ROOT_DIR, "i", "/mnt/cro.cz/Rundowns","Please specify the input path")
  flag.Parse()

  if len(ROOT_DIR) == 0{
    log.Fatal("Please specify the input folder -i ../..")
  }

  flag.Usage = func(){
    fmt.Println("Usage of program:")
    fmt.Println("./rundown_files_checker -i /path/to/annova/disk (full path to SMB mounted disk in your system)")
  }
}

func main() {


  // years list
  for _, year := range years {
    log.Println("Checking Rundown count Year "+year)
    files, err := ioutil.ReadDir(ROOT_DIR + "/" + year)
    if err != nil {
      log.Fatal(err)
    }

    // week list
    for _, f := range files {

      checked := 0
      var errornous_filenames []string

      files, err := ioutil.ReadDir(ROOT_DIR + "/" + year + "/" + f.Name())
      // files list
      for _, fn := range files{
        if strings.Contains(fn.Name(),"xml"){
          // unixtime := fmt.Sprint(xml.ModTime().Unix())
          week_no, _ := strconv.Atoi(strings.Split(fmt.Sprint(fn.ModTime().ISOWeek())," ")[1])
          dir_no, _ := strconv.Atoi( strings.Split(fmt.Sprint(f.Name()),"W")[1] )

          if week_no == dir_no{
            checked += 1
          }else{
            errornous_filenames = append(errornous_filenames, f.Name() + "/" + fn.Name())
          }

          // log.Printf("file was modded on: %v and is in dir %v (%s)",week_no,dir_no,fn.Name())
        }else{
            errornous_filenames = append(errornous_filenames, f.Name() + "/" + fn.Name())
        }
      }
      if err != nil{
        log.Fatal(err)
      }

      if(checked == len(files)){
        fmt.Println(year+"/"+f.Name() + ": comparing file modtime to foldername: " + fmt.Sprint(checked) + "/" + fmt.Sprint(len(files)) + " PASSED!")
      }else{
        fmt.Println(year+"/"+f.Name() + ": comparing file modtime to foldername: " + fmt.Sprint(checked) + "/" + fmt.Sprint(len(files)) + " NOT PASSED!")
        for _, ef := range errornous_filenames{
          fmt.Println("mismatch found: " + fmt.Sprint(ef))
        }
      }
    }
  }
}
