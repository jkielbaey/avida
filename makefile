SHELL := /bin/bash

BIN_DIR := $(GOPATH)/bin
GOMETALINTER := $(BIN_DIR)/gometalinter
PLATFORMS := windows linux darwin
BINARY := avida

# These will be provided to the target
VERSION := 0.0.1
BUILD := `git rev-parse HEAD`

# Use linker flags to provide version/build settings to the target
LDFLAGS=-ldflags "-X=main.version=$(VERSION) -X=main.build=$(BUILD)"

os = $(word 1, $@)

PKGS := $(shell go list ./... | grep -v /vendor)

$(GOMETALINTER):
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install &> /dev/null

clean:
	go clean
	rm -f $(BINARY)
	rm -f $(BINARY).exe
	rm -rf release/
.PHONY: clean

build:
	go build $(LDFLAGS) cmd/$(BINARY)/$(BINARY).go
.PHONY: build

test:
	go test $(PKGS)
.PHONY: test

lint: $(GOMETALINTER)
	gometalinter --vendor --config gometalinter.json  ./...
.PHONY: lint

$(PLATFORMS):
	mkdir -p release
	GOOS=$(os) GOARCH=amd64 go build $(LDFLAGS) -o release/$(BINARY)-v$(VERSION)-$(os)-amd64 cmd/$(BINARY)/$(BINARY).go
.PHONY: $(PLATFORMS)

install:
	@go install $(LDFLAGS) cmd/$(BINARY)/$(BINARY).go

# run "make release -j3
release: windows linux darwin
.PHONY: release
