package main

import (
  "fmt"
  "testing"
  "testing/fstest"
)


func TestMemoryFile(t *testing.T){
  fs := fstest.MapFS{
      "hello.txt": {
         Data: []byte("hello, world"),
      },
   }
   data, err := fs.ReadFile("hello.txt")
   if err != nil {
      t.Errorf("%q",err.Error())
   }
   t.Log(string(data) == "hello, world")
}


func TestCreateMessage(t *testing.T){

      message := Message{
			Index:  0,
			Status: "SUCCESS",
			Action: "none",
			Data: Data{
				Date: fmt.Sprintf("%04d-%02d-%02d", 2022, 1, 1),
				Week: fmt.Sprintf("W%02d", 1),
				File: "testName.xml",
			},
		}

      t.Logf("Message created: %q", message)
}

func TestSomething(t *testing.T){

  A := 0
  B := 0

  if(A != B){
    t.Errorf("Output %q not equal to expected %q", A, B)
  }
}

