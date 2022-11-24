package main

import (
	"fmt"
	//	"io/fs"
	//        "io/ioutil"
	//        "bytes"
	"testing"
	"testing/fstest"
)

// Test that mock file is created in memory
func TestMemoryFile(t *testing.T) {
	fs := fstest.MapFS{
		"hello.xml": {
			Data: mockData,
		},
	}
	fileHandle, err := fs.ReadFile("hello.xml")
	if err != nil {
		t.Errorf("%q", err.Error())
	}
	t.Log(len(string(fileHandle)) == 1383)
}

/*
// TODO: reading memory file to io.Reader (mock it here)

// here's a fake ReadFile method that matches the signature of ioutil.ReadFile
func ReadFile() (io.Reader, error) {
    buf := bytes.NewBufferString(string(mockData))
    return buf
}

func TestParseRundown(t *testing.T){

  year, month, day, week := ParseRundown(ReadFile())
  t.Logf("%q %q %q %q",year, month, day, week)

}

*/

// Test that Message struct is created
func TestCreateMessage(t *testing.T) {

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

	t.Logf("Message created: %q", message)
}

// Test that something equals
func TestSomething(t *testing.T) {

	A := 0
	B := 0

	if A != B {
		t.Errorf("Output %q not equal to expected %q", A, B)
	}
}

// --------------------------------------------------------------------
// DATA MOCK
// --------------------------------------------------------------------

// mocking actual rundown file data
var mockData []byte = []byte(`
<OPENMEDIA>
<OM_SERVER/>
<OM_OBJECT SystemID="3fc88f5c-ef6b-44fa-bdef-002c69855f16" ObjectID="0000000200957e65" DocumentURN="urn:openmedia:3fc88f5c-ef6b-44fa-bdef-002c69855f16:0000000200957E65" DirectoryID="0000000200003f57" InternalType="1" TemplateID="fffffffa00001022" TemplateType="1" TemplateName="Radio Rundown">
<OM_HEADER>
<OM_FIELD FieldID="1" FieldType="3" FieldName="Čas vytvoření" IsEmpty="no">
<OM_DATETIME>20211223T010314,000</OM_DATETIME>
</OM_FIELD>
<OM_FIELD FieldID="2" FieldType="3" FieldName="Aktualizováno kdy" IsEmpty="no">
<OM_DATETIME>20211223T010314,000</OM_DATETIME>
</OM_FIELD>
<OM_FIELD FieldID="3" FieldType="1" FieldName="Owner Name" IsEmpty="no">
<OM_STRING>admin</OM_STRING>
</OM_FIELD>
<OM_FIELD FieldID="5" FieldType="1" FieldName="Vytvořil" IsEmpty="no">
<OM_STRING>user_superuser</OM_STRING>
</OM_FIELD>
<OM_FIELD FieldID="6" FieldType="1" FieldName="Autor" IsEmpty="no">
<OM_STRING>user_superuser</OM_STRING>
</OM_FIELD>
<OM_FIELD FieldID="7" FieldType="1" FieldName="Titul" IsEmpty="yes">
<OM_STRING/>
</OM_FIELD>
<OM_FIELD FieldID="8" FieldType="1" FieldName="Název" IsEmpty="no">
<OM_STRING>00-05 ČRo Region SC - Čtvrtek 06.01.2022</OM_STRING>
</OM_FIELD>
<OM_FIELD FieldID="1004" FieldType="3" FieldName="Čas začátku" IsEmpty="no">
<OM_DATETIME>20220106T000000,000</OM_DATETIME>
</OM_FIELD>
</OM_HEADER>
</OM_OBJECT>
</OPENMEDIA>
`)
