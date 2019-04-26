SHELL       = /bin/bash -euo pipefail
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
	@goimports -ungroup -w .

.PHONY: generate
generate:
	@go generate ./...

.PHONY: refresh
refresh: generate format


.PHONY: test
test:
	@go test -race -timeout 1s ./...
