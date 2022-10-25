.PHONY: fmt clean build

.DEFAULT_GOAL := build`

BINARY_NAME=openmedia_checker

fmt:
	go fmt .
deps:
	go mod tidy

build: clean deps fmt
	go build -o ${BINARY_NAME} .

clean: 
	go clean
	rm ${BINARY_NAME}

run: clean deps fmt build
	./${BINARY_NAME}

