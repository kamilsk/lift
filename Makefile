# sourced by https://github.com/octomation/makefiles

.DEFAULT_GOAL = test-with-coverage

SHELL = /bin/bash -euo pipefail

GO          = GOPRIVATE=$(GOPRIVATE) GOFLAGS=$(GOFLAGS) go
GO111MODULE = on
GOFLAGS     = -mod=vendor
GOPRIVATE   = go.octolab.net
GOPROXY     = direct
LOCAL       = $(MODULE)
MODULE      = $(shell $(GO) list -m)
PACKAGES    = $(shell $(GO) list ./... 2> /dev/null)
PATHS       = $(shell $(GO) list ./... 2> /dev/null | sed -e "s|$(MODULE)/\{0,1\}||g")
TIMEOUT     = 1s

ifeq (, $(PACKAGES))
	PACKAGES = $(MODULE)
endif

ifeq (, $(PATHS))
	PATHS = .
endif

.PHONY: go-env
go-env:
	@echo "GO111MODULE: $(shell $(GO) env GO111MODULE)"
	@echo "GOFLAGS:     $(strip $(shell $(GO) env GOFLAGS))"
	@echo "GOPRIVATE:   $(strip $(shell $(GO) env GOPRIVATE))"
	@echo "GOPROXY:     $(strip $(shell $(GO) env GOPROXY))"
	@echo "LOCAL:       $(LOCAL)"
	@echo "MODULE:      $(MODULE)"
	@echo "PACKAGES:    $(PACKAGES)"
	@echo "PATHS:       $(strip $(PATHS))"
	@echo "TIMEOUT:     $(TIMEOUT)"

.PHONY: deps
deps:
	@$(GO) mod tidy
	@if [[ "$(shell $(GO) env GOFLAGS)" =~ -mod=vendor ]]; then $(GO) mod vendor; fi

.PHONY: deps-clean
deps-clean:
	@$(GO) clean -modcache

.PHONY: update
update: selector = '.Require[] | select(.Indirect != true) | .Path'
update:
	@if command -v egg > /dev/null; then \
		packages="$(shell egg deps list)"; \
		$(GO) get -mod= -u $$packages; \
	elif command -v jq > /dev/null; then \
		packages="$(shell $(GO) mod edit -json | jq -r $(selector))"; \
		$(GO) get -mod= -u $$packages; \
	else \
		packages="$(shell cat go.mod | grep -v '// indirect' | grep '\t' | awk '{print $$1}')"; \
		$(GO) get -mod= -u $$packages; \
	fi

BINARY  = $(BINPATH)/$(shell basename $(MODULE))
BINPATH = $(PWD)/bin
COMMIT  = $(shell git rev-parse --verify HEAD)
DATE    = $(shell date +%Y-%m-%dT%T%Z)
LDFLAGS = -ldflags "-s -w -X main.commit=$(COMMIT) -X main.date=$(DATE)"

export PATH := $(BINPATH):$(PATH)

.PHONY: build-env
build-env:
	@echo "BINARY:      $(BINARY)"
	@echo "BINPATH:     $(BINPATH)"
	@echo "COMMIT:      $(COMMIT)"
	@echo "DATE:        $(DATE)"
	@echo "LDFLAGS:     $(LDFLAGS)"

.PHONY: build
build: MAIN = .
build:
	@$(GO) build -o $(BINARY) $(LDFLAGS) $(MAIN)

.PHONY: build-clean
build-clean:
	@$(GO) clean -cache

.PHONY: install
install: build

.PHONY: install-clean
install-clean:
	@rm -f $(BINARY)

.PHONY: test
test:
	@$(GO) test -race -timeout $(TIMEOUT) $(PACKAGES)

.PHONY: test-clean
test-clean:
	@$(GO) clean -testcache

.PHONY: test-with-coverage
test-with-coverage:
	@$(GO) test -cover -timeout $(TIMEOUT) $(PACKAGES) | column -t | sort -r

.PHONY: test-with-coverage-profile
test-with-coverage-profile:
	@$(GO) test -cover -covermode count -coverprofile c.out -timeout $(TIMEOUT) $(PACKAGES)

.PHONY: dist
dist:
	@godownloader .goreleaser.yml > .github/install.sh

.PHONY: format
format:
	@goimports -local $(LOCAL) -ungroup -w $(PATHS)

.PHONY: generate
generate:
	@$(GO) generate $(PACKAGES)


.PHONY: clean
clean: build-clean deps-clean install-clean test-clean

.PHONY: env
env: go-env build-env

.PHONY: refresh
refresh: update deps generate format test build
