PKGS = $$(go list ./... | grep -v /vendor/)
PREFIX = $(shell pwd)
BUILDDIR = $(shell pwd)/bin

vet:
	@echo ">> vetting code"
	go vet $(PKGS)

fmt:
	@echo ">> fmt code"
	go fmt $(PKGS)

build:
	@echo ">> building code"
	go build -mod=vendor -tags=jsoniter -ldflags='-w -s -linkmode=external' -o $(BUILDDIR)/redisearchd $(PREFIX)/main.go

clean:
	@echo ">> clean build"
	rm -rf $(BUILDDIR)

api-doc:
	@echo "generating api doc..."
	swag init http/router.go
	@echo "Done."
