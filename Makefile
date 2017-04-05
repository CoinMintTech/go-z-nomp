default: build

build: build_darwin

build_linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s" -a -installsuffix cgo -o main.linux

build_darwin:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "-s" -a -installsuffix cgo -o main

vendor:
	glide install
