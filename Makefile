.PHONY: fmt clean build

.DEFAULT_GOAL := build`

BINARY_NAME=openmedia-files-checker

fmt:
	go fmt .
deps:
	go mod tidy

build: clean deps fmt
	go build -o ${BINARY_NAME} .

clean: 
	go clean

run: clean deps fmt build
	./${BINARY_NAME}

