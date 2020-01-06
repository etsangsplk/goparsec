.DEFAULT_GOAL := noop

GOBUILDFLAGS :=-a -v
GIT_COMMIT_TAG := $(shell git rev-parse --verify HEAD)
GIT_BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
PLATFORM := `uname | tr '[:upper:]' '[:lower:]'`

.phony: build test lint coverage

build:
	@go build $(GOBUILDFLAGS) ./...

test:
	@go test -v -race -timeout 4000s -test.run=. -test.bench=. -test.benchmem=true ./...

lint:
	@

coverage:
	@go test -coverprofile=coverage.out
	@go tool cover -html=coverage.out
	@rm -rf coverage.out

noop:
	@echo "No make target specified"

install-lint: VERSION = 1.21.0
install-lint:
	@echo "Installing golangci-lint...into $(GOPATH)/bin"
	curl -L https://github.com/golangci/golangci-lint/releases/download/v$(VERSION)/golangci-lint-$(VERSION)-$(PLATFORM)-amd64.tar.gz \
	| tar -xz -C $(GOPATH)/bin/ --strip-components 1 && rm $(GOPATH)/bin/README.md $(GOPATH)/bin/LICENSE
