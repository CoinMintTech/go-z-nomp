SHELL=/bin/bash
PLATFORM=$(shell go env GOOS)
ARCH=$(shell go env GOARCH)
GOPATH=$(shell go env GOPATH)
GOBIN=$(GOPATH)/bin

default: build

build:
	go fmt ./...
	DEP_BUILD_PLATFORMS=$(PLATFORM) DEP_BUILD_ARCHS=$(ARCH) ./bin/build-all.bash

test:
	go test ./...

vendor:
	dep ensure

.PHONY: build test
