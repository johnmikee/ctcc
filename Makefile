all: build

.PHONY: build

ifeq ($(GOPATH),)
	PATH := $(HOME)/go/bin:$(PATH)
else
	PATH := $(GOPATH)/bin:$(PATH)
endif

export GO111MODULE=on

BRANCH = $(shell git rev-parse --abbrev-ref HEAD)
REVISION = $(shell git rev-parse HEAD)
REVSHORT = $(shell git rev-parse --short HEAD)
USER = $(shell whoami)
GOVERSION = $(shell go version | awk '{print $$3}')
NOW	= $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
SHELL = /bin/sh
VERSION = $(shell git describe --tags --always)

ifneq ($(OS), Windows_NT)
	CURRENT_PLATFORM = linux
	ifeq ($(shell uname), Darwin)
		SHELL := /bin/sh
		CURRENT_PLATFORM = darwin
	endif
else
	CURRENT_PLATFORM = windows
endif

BUILD_VERSION = "\
	-X github.com/johnmikee/ctcc/version.appName=${APP_NAME} \
	-X github.com/johnmikee/ctcc/version.version=${VERSION} \
	-X github.com/johnmikee/ctcc/version.branch=${BRANCH} \
	-X github.com/johnmikee/ctcc/version.buildUser=${USER} \
	-X github.com/johnmikee/ctcc/version.buildDate=${NOW} \
	-X github.com/johnmikee/ctcc/version.revision=${REVISION} \
	-X github.com/johnmikee/ctcc/version.goVersion=${GOVERSION}"

deps:
	go mod download

test:
	go test -cover ./...

build: ctcc

clean:
	rm -rf build/
	rm -f *.zip

.pre-build:
	mkdir -p build/darwin

APP_NAME = ctcc

.pre-ctcc:
	$(eval APP_NAME = ctcc)

ctcc: .pre-build .pre-ctcc
	go build -o build/$(CURRENT_PLATFORM)/ctcc -trimpath -ldflags  ${BUILD_VERSION} .

package: ctcc
	mkdir -p package/payload/usr/local/bin/
	mv build/$(CURRENT_PLATFORM)/ctcc package/payload/usr/local/bin/
	cd package && munkipkg .

install: .pre-ctcc
	go install -trimpath -ldflags ${BUILD_VERSION} .

static-check:
	staticcheck ./...

vet:
	go vet ./...

tidy:
	go mod tidy
