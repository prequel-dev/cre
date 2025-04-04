SHELL := /bin/bash

VERSION ?= $(shell head -n 1 VERSION 2> /dev/null || echo "0.0.0")
GITHASH := $(shell git rev-parse HEAD)

LD_FLAGS_GITHASH := -X 'github.com/prequel-dev/cre/pkg/ruler.Githash=$(GITHASH)'
LD_FLAGS_VERSION := -X 'github.com/prequel-dev/cre/pkg/ruler.Version=$(VERSION)'
LD_FLAGS := $(LD_FLAGS_GITHASH) $(LD_FLAGS_VERSION) 

.PHONY: all
all: clean ruler rules

.PHONY: ruler
ruler:
	@mkdir -p bin/
	@env CGO_ENABLED=0 go build -ldflags "${LD_FLAGS}" -o ./bin/ruler ./cmd/ruler/ruler.go

.PHONY: rules
rules:
	@mkdir -p bin/rules/
	@./bin/ruler build -p rules  -o ./bin

.PHONY: clean
clean:
	rm -rf bin/*