SHELL       = /bin/bash -euo pipefail
PKGS        = go list ./... | grep -v vendor
GO111MODULE = on
GOFLAGS     = -mod=vendor
TIMEOUT     = 1s
BIN         = $(shell basename $(shell pwd))


.DEFAULT_GOAL = test-with-coverage


.PHONY: deps
deps:
	@go mod tidy && go mod vendor && go mod verify

.PHONY: update
update:
	@go get -mod= -u


.PHONY: format
format:
	@goimports -local $(dirname $(go list -m)) -ungroup -w .

.PHONY: generate
generate:
	@go generate ./...

.PHONY: refresh
refresh: generate format


.PHONY: test
test:
	@go test -race -timeout $(TIMEOUT) ./...

.PHONY: test-with-coverage
test-with-coverage:
	@go test -cover -timeout $(TIMEOUT) ./... | column -t | sort -r

.PHONY: test-with-coverage-profile
test-with-coverage-profile:
	@go test -cover -covermode count -coverprofile c.out -timeout $(TIMEOUT) ./...


.PHONY: sync
sync:
	@git stash && git pull --rebase && git stash pop || true

.PHONY: upgrade
upgrade: sync update deps refresh test-with-coverage


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
	@go run main.go -f testdata/app.toml up -m 6379:16379 -m 5672:15672 -m 5432:15432
	@echo ---
	@go run main.go -f testdata/app.toml down -m 6379:16379 -m 5672:15672 -m 5432:15432
	@echo ---
	@go run main.go -f testdata/app.toml env -m 6379:16379 -m 5672:15672 -m 5432:15432
	@echo ---
	@go run main.go -f testdata/app.toml forward -m 6379:16379 -m 5672:15672 -m 5432:15432
	@echo ---
	@go run main.go -f testdata/app.toml call -m 6379:16379 -m 5672:15672 -m 5432:15432 -- echo '$$REDIS_PORT $$PGPORT'
