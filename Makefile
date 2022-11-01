.PHONY: fmt clean build
.DEFAULT_GOAL := build

BINARY_NAME=openmedia-files-checker

fmt:
	go fmt .
deps:
	go mod tidy
clean: 
	go clean

build: clean deps
	go build -o ${BINARY_NAME} .

run: fmt clean deps build
	./${BINARY_NAME}

install: fmt clean deps build
	cp ./${BINARY_NAME} /usr/local/bin/

uninstall: clean
	rm /usr/local/bin/${BINARY_NAME}
