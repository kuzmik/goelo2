SHELL := /bin/bash

# The name of the executable (default is current directory name)
TARGET := $(shell echo $${PWD\#\#*/})

# These will be provided to the target
VERSION := 0.1.0
BUILD := `git rev-parse HEAD`
BUILD_TIME=`date +%FT%T%z`

# Use linker flags to provide version/build settings to the target.
# If we need debugging symbols, remove -s and -w
LDFLAGS=-ldflags "-s -w -X=main.Version=$(VERSION) -X=main.Build=$(BUILD) -X=main.BuildTime=$BUILD_TIME)"
LINUX_LDFLAGS=-ldflags "-s -w -X=main.Version=$(VERSION) -X=main.Build=$(BUILD) -X=main.BuildTime=$BUILD_TIME) -linkmode external -extldflags -static"

# go source files, ignore vendor directory
SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")

all: clean build

$(TARGET): $(SRC)
	@go build $(LDFLAGS) -o $(TARGET)

build: clean $(TARGET)
	@true

run: build
	@./$(TARGET)

clean:
	@rm -f $(TARGET)

check:
	@gofmt -s -l -w .
	@golint .
	@go vet ./...

# cross compile for linux
# brew install FiloSottile/musl-cross/musl-cross
linux: clean $(TARGET)
	@CC=x86_64-linux-musl-gcc CXX=x86_64-linux-musl-g++ GOARCH=amd64 GOOS=linux CGO_ENABLED=1 go build $(LINUX_LDFLAGS) -o $(TARGET)

# brew install upx
strip:
	@upx --brute $(TARGET)
