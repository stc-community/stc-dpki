.PHONY: all build clean

PROG=bin/dpki
SRCS=cmd/main.go

# git commit hash
COMMIT_HASH=$(shell git rev-parse --short HEAD || echo "GitNotFound")
# Compilation date
BUILD_DATE=$(shell date '+%Y-%m-%d %H:%M:%S')
# Compilation conditions
CFLAGS = -ldflags "-s -w -X \"main.BuildVersion=${COMMIT_HASH}\" -X \"main.BuildDate=$(BUILD_DATE)\""

all:
	if [ ! -d "./bin/" ]; then \
	mkdir bin; \
	fi
	CGO_ENABLED=0 go build  $(CFLAGS) -o $(PROG) $(SRCS)

release:
	if [ ! -d "./bin/" ]; then \
	mkdir bin; \
	fi
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  $(CFLAGS) -o $(PROG) $(SRCS)

wire:
	@wire gen ./internal

clean:
	rm -rf ./bin
	rm -rf ./data