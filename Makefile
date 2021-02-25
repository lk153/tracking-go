.PHONY: all one two

GOPATH := $(or $(GOPATH), $(HOME)/go)
DIST_DIR := out
WIRE := $(GOPATH)/bin/wire

$(WIRE):
	GOPATH=$(GOPATH) go install -mod=mod github.com/google/wire/cmd/wire

default: di build;

di: $(WIRE)
	$(WIRE) gen -output_file_prefix build_server_  ./cmd

build: 
	go build -o $(DIST_DIR) ./cmd