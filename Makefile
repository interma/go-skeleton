all: build test

SHELL := /bin/bash
BINARY = main
BIN_DIR = $(shell echo $${GOPATH:-~/go} | awk -F':' '{ print $$1 "/bin"}')

GOFLAGS :=
VERSION_STR = ""
#VERSION_STR = "-X github.com/greenplum-db/gpmigrate/migrate.version=$(shell date +%Y-%m-%d)"

.PHONY: depend format unit coverage build build_mac build_linux clean

depend:
	go get golang.org/x/tools/cmd/goimports
	go get github.com/golang/lint/golint
	go get github.com/onsi/ginkgo/ginkgo
	go get github.com/alecthomas/gometalinter
	gometalinter --install
	go get github.com/golang/dep/cmd/dep
	dep ensure

format:
	gofmt -w -s main.go pkg
	goimports -w main.go pkg

lint:
	gometalinter --config=gometalinter.config -s vendor ./...

unit: format
	ginkgo -r -cover -coverprofile=coverage.out pkg/add

coverage: unit
	go tool cover -func=pkg/add/coverage.out

#end2end:
#	ginkgo -r end2end

build: format
	go build -tags '$(BINARY)' $(GOFLAGS) -o $(BIN_DIR)/$(BINARY) -ldflags $(VERSION_STR)

build_mac: format
	env GOOS=darwin GOARCH=amd64 go build -tags '$(BINARY)' $(GOFLAGS) -o $(BINARY)_mac -ldflags $(VERSION_STR)

build_linux: format
	env GOOS=linux GOARCH=amd64 go build -tags '$(BINARY)' $(GOFLAGS) -o $(BINARY)_linux -ldflags $(VERSION_STR)

clean:
	rm -f $(BINARY)_mac $(BINARY)_linux $(BIN_DIR)/$(BINARY) pkg/add/coverage.out
