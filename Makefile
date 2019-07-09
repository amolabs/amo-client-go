.PHONY: build test

all: build

GO := $(shell command -v go 2> /dev/null)
FS := /
# go source code files including files from vendor directory
GOSRCS=$(shell find . -name \*.go)
#BUILDENV=CGO_ENABLED=0

ifeq ($(GO),)
  $(error could not find go. Is it in PATH? $(GO))
endif

ifneq ($(TARGET),)
  BUILDENV += GOOS=$(TARGET)
endif

GOPATH ?= $(shell $(GO) env GOPATH)
GITHUBDIR := $(GOPATH)$(FS)src$(FS)github.com

GOPATH ?= $(shell $(GO) env GOPATH)

go_get = $(if $(findstring Windows_NT,$(OS)),\
IF NOT EXIST $(GITHUBDIR)$(FS)$(1)$(FS) ( mkdir $(GITHUBDIR)$(FS)$(1) ) else (cd .) &\
IF NOT EXIST $(GITHUBDIR)$(FS)$(1)$(FS)$(2)$(FS) ( cd $(GITHUBDIR)$(FS)$(1) && git clone https://github.com/$(1)/$(2) ) else (cd .) &\
,\
mkdir -p $(GITHUBDIR)$(FS)$(1) &&\
(test ! -d $(GITHUBDIR)$(FS)$(1)$(FS)$(2) && cd $(GITHUBDIR)$(FS)$(1) && git clone https://github.com/$(1)/$(2)) || true &&\
)\
cd $(GITHUBDIR)$(FS)$(1)$(FS)$(2) && git fetch origin && git checkout -q $(3)

tags: $(GOSRCS)
	gotags -R -f tags .

get_tools: $(GOPATH)/bin/dep

$(GOPATH)/bin/dep:
	$(call go_get,golang,dep,22125cfaa6ddc71e145b1535d4b7ee9744fefff2)
	cd $(GITHUBDIR)$(FS)golang$(FS)dep$(FS)cmd$(FS)dep && $(GO) install

get_vendor_deps:
	@echo "--> Generating vendor directory via dep ensure"
	@rm -rf .vendor-new
	@dep ensure -v -vendor-only

update_vendor_deps:
	@echo "--> Running dep ensure"
	@rm -rf .vendor-new
	@dep ensure -v -update

build:
	@echo "--> Building amo console (amocli)"
	$(BUILDENV) go build ./cmd/amocli

install:
	@echo "--> Installing amo console (amocli)"
	go install ./cmd/amocli

test:
	go test ./...

clean:
	rm -f amocli
