SHELL := /bin/bash

VERSION_FILE := VERSION
SEMVER := $(shell cat $(VERSION_FILE))

VERSION ?= $(shell head -n 1 VERSION 2> /dev/null || echo "0.0.0")
GITHASH := $(shell git rev-parse HEAD)

LD_FLAGS_GITHASH := -X 'github.com/prequel-dev/cre/pkg/ruler.Githash=$(GITHASH)'
LD_FLAGS_VERSION := -X 'github.com/prequel-dev/cre/pkg/ruler.Version=$(VERSION)'
LD_FLAGS := $(LD_FLAGS_GITHASH) $(LD_FLAGS_VERSION) 

BINDIR := ./bin

.PHONY: all
all: clean ruler rules

.PHONY: bindir
bindir:
	mkdir -p $(BINDIR)

.PHONY: ruler
ruler: bindir
	@env CGO_ENABLED=0 go build -ldflags "${LD_FLAGS}" -o ./bin/ruler ./cmd/ruler/ruler.go

.PHONY: linux
linux: bindir
	@env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -buildmode=pie -trimpath -ldflags "${LD_FLAGS}" -o ./bin/ruler-linux-amd64 ./cmd/ruler/ruler.go
	@env GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -buildmode=pie -trimpath -ldflags "${LD_FLAGS}" -o ./bin/ruler-linux-arm64 ./cmd/ruler/ruler.go

.PHONY: darwin
darwin: bindir
	@env GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 go build -buildmode=pie -trimpath -ldflags "${LD_FLAGS}" -o ./bin/ruler-darwin-arm64 ./cmd/ruler/ruler.go

.PHONY: windows
windows: bindir
	@env GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -buildmode=pie -trimpath -ldflags "${LD_FLAGS}" -o ./bin/ruler-windows-amd64 ./cmd/ruler/ruler.go

.PHONY: rules
rules: ruler
	@mkdir -p bin/rules/
	@./bin/ruler build -p rules -o ./bin

.PHONY: clean
clean:
	@rm -rf bin/*

.PHONY: test
test:
	@go test ./...
