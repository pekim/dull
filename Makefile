SHELL := /bin/bash

TAGS =

SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")
MAIN_PKG = github.com/pekim/dull3

fmt:
	@gofmt -l -w $(SRC)

install: install_tools dep_ensure install_packages

install_packages:
	@go install $(TAGS) $(MAIN_PKG)

install_tools:
	@go get -v -u github.com/golang/dep/...
	@go get -v -u github.com/jteeuwen/go-bindata/...

dep_ensure:
	@dep ensure

bindata:
	go generate

run_simple:
	go run _demo/simple/main.go
