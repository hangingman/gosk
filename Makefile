# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOINSTALL=$(GOCMD) install
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

BIN=gosk
NASK=wine nask.exe

.PHONY: all testdata

all: dep test build

build:
	cd cmd/gosk && $(GOBUILD) -v
	cd cmd/f12copy && $(GOBUILD) -v
	cd cmd/f12format && $(GOBUILD) -v
	cd ..
	$(GOINSTALL) -v ./...

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)

run: build
	./$(BIN)

fmt:
	for go_file in `find . -name \*.go`; do \
		go fmt $${go_file}; \
	done

dep:
	$(GOGET) github.com/stretchr/testify
	$(GOGET) github.com/comail/colog
	$(GOGET) github.com/pk-rawat/gostr
	$(GOGET) github.com/mitchellh/go-fs

emacs:
	$(GOGET) -u github.com/rogpeppe/godef
	$(GOGET) -u github.com/nsf/gocode
	$(GOGET) -u golang.org/x/lint/golint
	$(GOGET) -u github.com/kisielk/errcheck
	$(GOGET) -u github.com/derekparker/delve/cmd/dlv

testdata:
	$(NASK) testdata/byte-opcode.nas testdata/byte-opcode.obj testdata/byte-opcode.list
