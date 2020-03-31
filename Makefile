PROJECT_NAME := "money"
PKG := "github.com/radicalcompany/$(PROJECT_NAME)"
ENVGO := $(shell go env GOPATH)
.PHONY: all swag dep dep-test dep-build build clean test coverage coverhtml lint install-lint

all: build

install-lint: ## install golint
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(ENVGO)/bin v1.23.6

lint: ## Lint the files
	GL_DEBUG=linters_output;GOPACKAGESPRINTGOLISTERRORS=1;golangci-lint run -E goimports -E misspell -E whitespace -E goprintffuncname -E godox

test: ## Run unittests
	@go test -v

dep: dep-test dep-build

dep-test:
	@go get -t

dep-build:
	@go get

build: dep ## Build the binary file
	@go build -i -v $(PKG)

clean: ## Remove previous build
	@rm -f $(PROJECT_NAME)

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
