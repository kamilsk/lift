.DEFAULT_GOAL = test-with-coverage

SHELL = /bin/bash -euo pipefail

GO111MODULE = on
GOFLAGS     = -mod=vendor
MODULE      = $(shell go list -m)
PACKAGES    = $(shell go list ./...)
PATHS       = $(shell go list ./... | sed -e "s|$(shell go list -m)/\{0,1\}||g")
TIMEOUT     = 1s
VENDOR      = $(dir $(MODULE))

.PHONY: go-env
go-env:
	@echo "GO111MODULE: $(shell go env GO111MODULE)"
	@echo "GOFLAGS:     $(strip $(shell go env GOFLAGS))"
	@echo "MODULE:      $(MODULE)"
	@echo "PACKAGES:    $(PACKAGES)"
	@echo "PATHS:       $(strip $(PATHS))"
	@echo "TIMEOUT:     $(TIMEOUT)"
	@echo "VENDOR:      $(VENDOR)"

BINPATH = $(PWD)/bin
BINARY  = $(BINPATH)/$(shell basename $(shell go list -m))
COMMIT  = $(shell git rev-parse --verify HEAD)
DATE    = $(shell date +%Y-%m-%dT%T%Z)

export PATH := $(BINPATH):$(PATH)

.PHONY: bin-env
bin-env:
	@echo "BINARY:  $(BINARY)"
	@echo "BINPATH: $(BINPATH)"
	@echo "COMMIT:  $(COMMIT)"
	@echo "DATE:    $(DATE)"

.PHONY: deps
deps:
	@go mod tidy && go mod vendor && go mod verify

.PHONY: update
update:
	@go get -mod= -u

.PHONY: format
format:
	@goimports -local $(VENDOR) -ungroup -w $(PATHS)

.PHONY: generate
generate:
	@go generate $(PACKAGES)

.PHONY: test
test:
	@go test -race -timeout $(TIMEOUT) $(PACKAGES)

.PHONY: test-with-coverage
test-with-coverage:
	@go test -cover -timeout $(TIMEOUT) $(PACKAGES) | column -t | sort -r

.PHONY: test-with-coverage-profile
test-with-coverage-profile:
	@go test -cover -covermode count -coverprofile c.out -timeout $(TIMEOUT) $(PACKAGES)

.PHONY: build
build:
	@go build -o $(BINARY) -ldflags "-s -w -X main.commit=$(COMMIT) -X main.date=$(DATE)" .

.PHONY: dist
dist:
	@godownloader .goreleaser.yml > .github/install.sh

