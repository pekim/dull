SHELL := /bin/bash

TARGET := dist/miked

# These will be provided to the target
VERSION := 1.0.0
BUILD := `git rev-parse HEAD`

# Use linker flags to provide version/build settings to the target
LDFLAGS=-ldflags "-X=main.Version=$(VERSION) -X=main.Build=$(BUILD)"

TAGS =

SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")
MAIN = cmd/miked/main.go
MAIN_PKG = github.com/pekim/miked/cmd/miked

all: build

$(TARGET): $(SRC)
	@go build $(LDFLAGS) -o $(TARGET) $(MAIN)

build: $(TARGET)
	@true

clean:
	@rm -f $(TARGET)

fmt:
	@gofmt -l -w $(SRC)

install: install_tools dep_ensure install_packages

install_packages:
	go install $(TAGS) $(MAIN_PKG)

install_tools:
	go get -v -u github.com/golang/dep/...

dep_ensure:
	dep ensure

run: build
	@$(TARGET)
