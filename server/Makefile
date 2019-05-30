HAVE_GOLINT:=$(shell which golint)

## Go
.PHONY: setup lint test build run
setup:
	@echo "Start setup"
	@env GO111MODULE=on go mod vendor

lint: setup golint
	@echo "Check lint"
	@golint $(shell go list ./...|grep -v vendor)
	@go vet ./...

test: setup
	@echo "go test"
	@go test

build: setup
	@echo "build"
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./bin/server .

run: setup
	@echo "go run"
	@go run ./main.go

## Install package
.PHONY: golint
golint:
ifndef HAVE_GOLINT
	@echo "Installing linter"
	@go get -u github.com/golang/lint/golint
endif
