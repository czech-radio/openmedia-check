package main

import (
	"fmt"
	"io"
	"strings"
	"testing"
	"testing/fstest"

	"github.com/spf13/afero"
)

// Test that mock file is created in memory
func TestMemoryFile(t *testing.T) {
         fs := fstest.MapFS{
		"hello1.xml": {
			Data: mockData,
		},
        "hello2.xml": {
			Data: mockData,
		},
            }
        fileHandle, err := fs.ReadFile("hello1.xml")
	if err != nil {
		t.Errorf("%q", err.Error())
	}

	t.Log(len(string(fileHandle)))
}

// afero mocking filesystem
func TestReportRundown(t *testing.T) {
    //virtual filesystem
    appFS := afero.NewMemMapFs()
    afero.WriteFile(appFS,"one.xml", mockData, 0644)
    afero.WriteFile(appFS,"two.xml", mockData, 0644)

    files, err := afero.ReadDir(appFS,"/")
    if err != nil{
        t.Error("Error opening file")
    }
    t.Log(files)
    
    
    //ReportRundowns("/","/",files,true)

}
// helper function
func ReadFile(bytes []byte) io.Reader {
        return strings.NewReader(string(bytes))
}

// test ParseRundown on mockData
func TestParseRundown(t *testing.T) {
        t.Logf("%T",ReadFile(mockData))
        // bytes, and utf-16 false
	year, month, day, week := ParseRundown(ReadFile(mockData), false)
        // returns 0,0,0,0 due to utf-16 reader object (probably)
        t.Logf("%v %v %v %v", year, month, day, week)
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

// Test that something equals
func TestSomething(t *testing.T) {

	A := 0
	B := 0

	if A != B {
		t.Errorf("Output %q not equal to expected %q", A, B)
                t.Fail()
	}
}

// --------------------------------------------------------------------
// DATA MOCK
// --------------------------------------------------------------------

// mocking actual rundown file data
var mockData []byte = []byte(`
<OPENMEDIA>
<OM_OBJECT SystemID="3fc88f5c-ef6b-44fa-bdef-002c69855f16" ObjectID="0000000200957e65" DocumentURN="urn:openmedia:3fc88f5c-ef6b-44fa-bdef-002c69855f16:0000000200957E65" DirectoryID="0000000200003f57" InternalType="1" TemplateID="fffffffa00001022" TemplateType="1" TemplateName="Radio Rundown">
<OM_FIELD FieldID="1004" FieldType="3" FieldName="Čas začátku" IsEmpty="no"><OM_DATETIME>20220106T000000,000</OM_DATETIME></OM_FIELD>
</OM_OBJECT>
</OPENMEDIA>
`)
