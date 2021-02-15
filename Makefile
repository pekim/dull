SHELL := /bin/bash

TAGS =

SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")
MAIN_PKG = github.com/pekim/dull

fmt:
	@gofmt -l -w $(SRC)

install: install_tools install_packages

install_packages:
	go install $(TAGS) $(MAIN_PKG)

test:
	@go test -v github.com/pekim/dull/...

run_simple:
	@go run _demo/simple/main.go
