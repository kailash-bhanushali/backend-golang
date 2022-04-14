.PHONY: FORCE

export GO111MODULE=on
export GOSUMDB=off
export GONOPROXY=github.com/kailash-bhanushali/backend-golang
export GOPROVATE=github.com/kailash-bhanushali/backend-golang
export GONOSUMDB=github.com/kailash-bhanushali/backend-golang

GO := go
BIN_DIR := $(GOPATH)/bin
VERSION := $(shell git describe --tags --always --dirty)
GOFILES := $(shell find . -name '*.go' -type f -not -name '*.pb.go' -not -name '*_generated.go' -not -name '*_test.go')
GOTESTS := $(shell find . -name '*.go' -type f -not -name '*.pb.go' -not -name '*_generated.go' -name '*_test.go')
BUILD_FILE_NAME := sample
MAIN_FILE_PATH := ./cmd

.PHONY: clean
clean:
	@rm -rf $(BIN_DIR)

.PHONY: get
get:
	$(GO) mod download
	$(GO) mod verify
	$(GO) mod tidy -compat=1.17

.PHONY: update
	$(GO) gte -u -v all

.PHONY: fmt
	gofmt -s -l -w $(GOFILES) $(GOTESTS)

.PHONY: build
build:
	@env GOARCH=amd64 GOOS=linux CGO_ENABLED=1 $(GO) build -o $(BIN_DIR)/$(BUILD_FILE_NAME) $(MAIN_FILE_PATH)

.PHONY: test
test:
	$(GO) test -mod=mod -race -coverprofile=coverage.out ./...
	$(GO) test cover -func coverage.out | tail -n 1 | awk '{print "Total Coverage: " $$3 }'

.PHONY: lint
lint:
	golint $(shell go list ./...)

.PHONY: doc
doc:
	$(GO) get github.com/swaggo/swag/cmd/swag
	$(BIN_DIR)/swag init

.PHONY: vendor
vendor:
	$(GO) mod vendor

.PHONY: swag-local
swag-local:
	swag init -g $(MAIN_FILE_PATH)/main.go