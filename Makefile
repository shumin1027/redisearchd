PWD=$(shell pwd)
DIST=$(shell pwd)/bin
DATE=$(shell date --iso-8601=seconds)

GIT_SHA=$(shell git rev-parse HEAD)
GIT_BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
GIT_CLOSEST_TAG=$(shell git describe --abbrev=0 --tags)

PKGS=$(shell go list ./... | grep -v /vendor/)
PROJECT="redisearchd"

BUILD_INFO_IMPORT_PATH=gitlab.xtc.home/xtc/redisearchd/app
BUILD_INFO='-X $(BUILD_INFO_IMPORT_PATH).BuildTime=$(DATE) -X $(BUILD_INFO_IMPORT_PATH).GitCommit=$(GIT_SHA) -X $(BUILD_INFO_IMPORT_PATH).GitBranch=$(GIT_BRANCH) -X $(BUILD_INFO_IMPORT_PATH).GitTag=$(GIT_CLOSEST_TAG)'

GOPATH := $(HOME)/go
PATH := $(GOPATH)/bin/:$(PATH)

.PHONY: dev-tools
dev-tools:
ifneq (,$(wildcard $(GOPATH)/bin/staticcheck))
	@staticcheck -version
else
	@go install honnef.co/go/tools/cmd/staticcheck@latest
endif
ifneq (,$(wildcard $(GOPATH)/bin/golangci-lint))
	@golangci-lint --version
else
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
endif
ifneq (,$(wildcard $(GOPATH)/bin/swag))
	@swag --version
else
	@go install github.com/swaggo/swag/cmd/swag@latest
endif

.PHONY: build
build:clean fmt vet check doc
	@echo ">> building code"
	go build -mod=vendor -tags=jsoniter -ldflags='-w -s -linkmode=external' -ldflags=$(BUILD_INFO) -o $(DIST)/redisearchd $(PWD)/main.go
	strip $(DIST)/redisearchd
	upx $(DIST)/redisearchd

.PHONY: fmt
fmt:
	@echo ">> fmt code"
	go fmt $(PKGS)

.PHONY: vet
vet:
	@echo ">> vetting code"
	go vet $(PKGS)

.PHONY: check
check:
	@echo ">> staticcheck code"
	staticcheck $(PKGS)

.PHONY: lint
lint:
	@echo ">> lint code"
	golangci-lint run

.PHONY: fix
fix:
	@echo ">> fix code"
	golangci-lint run --fix

.PHONY: clean
clean:
	@echo ">> clean build"
	go clean -i -x 
	rm -rf $(DIST)

.PHONY: clean-cache
clean-cache:
	@echo ">> clean build cache"
	go clean -cache -testcache

.PHONY: vendor
vendor:
	go mod tidy
	go mod vendor
	go mod verify

.PHONY: update
update:
	go get -u

.PHONY: doc
doc:
	@echo ">>Generating API DOC"
	swag init http/router.go
	@echo "Done."
