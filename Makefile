.PHONY: fmt clean build
.DEFAULT_GOAL := build

BINARY_NAME=openmedia_files_checker

fmt:
	go fmt .
deps:
	go mod tidy
clean: 
	go clean
	rm ./${BINARY_NAME}

build: deps
	go build -o ${BINARY_NAME} .

run: fmt clean deps build
	./${BINARY_NAME}

install: fmt clean deps build
	cp ./${BINARY_NAME} /usr/local/bin/

uninstall: clean
	rm /usr/local/bin/${BINARY_NAME}
