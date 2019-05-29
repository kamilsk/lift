SHELL       = /bin/bash -euo pipefail
BIN         = $(shell basename $(shell pwd))
PKGS        = go list ./... | grep -v vendor | grep -v ^_
GO111MODULE = on
GOFLAGS     = -mod=vendor


.PHONY: deps
deps:
	@go mod tidy && go mod vendor && go mod verify

.PHONY: update
update:
	@go get -mod= -u


.PHONY: format
format:
	@lift -f testdata/app.toml call -- goimports -local '$$GOMODULE' -ungroup -w .

.PHONY: generate
generate:
	@go generate ./...

.PHONY: refresh
refresh: generate format


.PHONY: test
test:
	@go test -race -timeout 1s ./...

.PHONY: test-with-coverage-profile
test-with-coverage-profile:
	@go test -covermode count -coverprofile c.out -timeout 1s ./...


.PHONY: build
build:
	@go build -o bin/$(BIN) .

.PHONY: dist
dist:
	@godownloader .goreleaser.yml > .github/install.sh

.PHONY: install
install:
	@go build -o $(GOPATH)/bin/$(BIN) .

.PHONY: run
run:
	@go run main.go env -f testdata/app.toml
	@echo ---
	@go run main.go forward -f testdata/app.toml
	@echo ---
	@go run main.go up -f testdata/app.toml -- main.go -f testdata/app.toml version
