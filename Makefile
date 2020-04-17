.PHONY: all test test-coverage

SOURCE_DIRS=$(shell go list ./... | grep -v /vendor | grep -v /out | cut -d "/" -f4  | uniq)
export GO111MODULE=on
BINARY_NAME=json_markd
APP_EXECUTABLE="out/json_markd"

SHELL := /bin/bash # Use bash syntax
default: check-quality test build

ALL_PACKAGES=$(shell go list ./... | grep -v /vendor)
WORKDIR=$(shell echo "${PWD}")
APPLICATION_YAML=$(shell echo "$(WORKDIR)/application.yml")
LOCALE_FILES_PATH=$(shell echo "${WORKDIR}/translation/locales"| sed 's:/:\\\/:g')

all: check-quality build test

check-quality: lint fmt vet

build:
	mkdir -p out/
	go build -o $(APP_EXECUTABLE)

fmt:
	go fmt ./...

vet:
	go vet ./...

lint:
	golint ./... | { grep -vwE "exported (var|function|method|type|const) \S+ should have comment" || true; }

test: 
	go clean -testcache ./... && go test -v ./... -coverprofile=coverage.out
	go tool cover -func coverage.out

test-coverage:
	go tool cover -html=coverage.out

clean: 
	go clean
	rm -f $(BINARY_NAME)

run:
	go build -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)
