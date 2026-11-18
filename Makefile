GO      ?= go
GOFMT   ?= gofmt
BIN     := bin/opsbot
VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo dev)
LDFLAGS := -X main.version=$(VERSION)

.PHONY: build test lint

## build: compile bin/opsbot with the version from git describe
build:
	$(GO) build -ldflags "$(LDFLAGS)" -o $(BIN) ./cmd/opsbot
	@./$(BIN) version

## test: run all unit tests
test:
	$(GO) test ./...

## lint: fail if code is not gofmt-formatted, then run go vet
lint:
	@out="$$($(GOFMT) -l .)"; if [ -n "$$out" ]; then echo "not gofmt-formatted:"; echo "$$out"; exit 1; fi
	$(GO) vet ./...
