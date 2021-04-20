PKGS = $$(go list ./... | grep -v /vendor/)
PREFIX = $(shell pwd)
BUILDDIR = $(shell pwd)/bin

.PHONY: build
build:clean fmt vet doc
	@echo ">> building code"
	go build -mod=vendor -tags=jsoniter -ldflags='-w -s -linkmode=external' -o $(BUILDDIR)/redisearchd $(PREFIX)/main.go
	strip $(BUILDDIR)/redisearchd

.PHONY: fmt
fmt:
	@echo ">> fmt code"
	go fmt $(PKGS)

.PHONY: vet
vet:
	@echo ">> vetting code"
	go vet $(PKGS)

.PHONY: clean
clean:
	@echo ">> clean build"
	go clean -i -x 
	rm -rf $(BUILDDIR)

.PHONY: clean-cache
clean-cache:
	@echo ">> clean build cache"
	go clean -cache -testcache

.PHONY: vendor
vendor:
	@go mod tidy
	@go mod verify
	@go mod vendor

.PHONY: update
update:
	@go get -u

.PHONY: doc
doc:
	@echo ">>Generating API DOC"
	swag init http/router.go
	@echo "Done."
