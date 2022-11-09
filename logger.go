package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type JSONLogger struct {
	Filename string
	Messages []Message
}

type Message struct {
	Date string `json:"date"`
	Type string `json:"type"`
	Msg  string `json:"message"`
	Status int `json:"status"`
}

func (logger *JSONLogger) Init(filename string) {
	logger.Filename = filename
	logger.Messages = make([]Message, 0)

}

func (logger *JSONLogger) Println(message string) (string, error) {
	messageType := "info"
	status := 0
        newMsg := Message{Date: time.Now().String(), Type: messageType, Status: status, Msg: message}
	logger.Messages = append(logger.Messages, newMsg)

	val, err := json.MarshalIndent(newMsg,"","    ")
	if err != nil {
		return "", err
	}
	fmt.Println(string(val))
	return string(val), nil
}

func (logger *JSONLogger) Fatal(message string) (string, error) {
	messageType := "error"
	status := 1
        newMsg := Message{Date: time.Now().String(), Type: messageType, Status: status, Msg: message}
	logger.Messages = append(logger.Messages, newMsg)

	val, err := json.MarshalIndent(newMsg,"","    ")
	if err != nil {
		return "", err
	}
	fmt.Println(string(val))
	return string(val), nil
}

func (logger *JSONLogger) Warn(message string) (string, error) {
	messageType := "warning"
	status := 2
        newMsg := Message{Date: time.Now().String(), Type: messageType, Status: status, Msg: message}
	logger.Messages = append(logger.Messages, newMsg)

	val, err := json.MarshalIndent(newMsg,"","    ")
	if err != nil {
		return "", err
	}
	fmt.Println(string(val))
	return string(val), nil
}

func (logger *JSONLogger) Write() {
	msg, err := json.MarshalIndent(logger.Messages, "", "    ")
	if err != nil {
		fmt.Println(err)
	}

	f, err := os.OpenFile(logger.Filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err = f.WriteString(string(msg)); err != nil {
		panic(err)
	}

}
