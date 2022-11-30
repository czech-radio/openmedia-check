package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"
)

var Messages []Message

// -----------------------------------------------------------------
// Test on real data
// -----------------------------------------------------------------

func TestReportRundown(t *testing.T) {

	path := filepath.Join("test", "data", "Rundowns", "2022", "W01")
	files, err := ioutil.ReadDir(path)
	if err != nil {
		t.Error("Error opening test folder")
	}
	t.Log(files)

	Messages = ReportRundowns(path, path, files)

}

// Test that Message struct is created and formatted right
func TestFormatMessage(t *testing.T) {
	message := Message{
		Index:  0,
		Status: "SUCCESS",
		Action: "none",
		Data: Data{
			Date: fmt.Sprintf("%04d-%02d-%02d", 2022, 1, 1),
			Week: fmt.Sprintf("W%02d", 1),
			File: "testName.xml",
			Dest: "data",
		},
	}

	FormatMessage(message)
}

// Test that move function is triggered
func TestRepairRundown(t *testing.T) {
	RepairRundows(Messages, true)
}
