.DEFAULT_GOAL := build

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

clean: vet
	go clean ./...

build: clean
	go build -o hello

.PHONY: go fmt vet clean build
