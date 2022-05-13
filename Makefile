PROJECT_NAME := "goatter"
PKG := "github.com/Dedalum/$(PROJECT_NAME)"
PKG_LIST := $(shell go list $(PKG)/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)

VERSIONSTRING := $(shell git describe --long --tags --always)
LDFLAGS := -X main.version=$(VERSIONSTRING)

ENV = $(ENVIRONMENT)

.PHONY: all lint test dep build dockerize clean

all: build

lint: ## Lint the files
		@golint -set_exit_status $(PKG_LIST)

test: ## Run unittests
		@CGO_ENABLED=0 GO111MODULE=off go test $(PKG_LIST)

dep: ## Get the dependencies
		@dep ensure

build: dep clean ## Build the binary file
		@CGO_ENABLED=0 GOOS=linux go build -a -ldflags "$(LDFLAGS)" -o build/$(PROJECT_NAME)/$(ENV)/$(PROJECT_NAME)_$(VERSIONSTRING) $(PKG)

clean: ## Remove previous build
		@rm -rf build

help: ## Display this help screen
		@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
		']]"}'
