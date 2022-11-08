package main

import (
	"encoding/json"
	"fmt"
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
}

func (logger *JSONLogger) Init(filename string) {
	logger.Filename = filename
	logger.Messages = make([]Message, 0)
}

func (logger JSONLogger) Println(messageType string, message string) (string, error) {

	newMsg := Message{Date: time.Now().String(), Type: messageType, Msg: message}
	logger.Messages = append(logger.Messages, newMsg)

	val, err := json.Marshal(newMsg)
	if err != nil {
		return "", err
	}
	fmt.Println(string(val))
	return string(val), nil
}

func (logger JSONLogger) list() {
	msg, err := json.Marshal(logger.Messages)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(msg))

}
