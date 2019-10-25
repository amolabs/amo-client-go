.PHONY: build test

all: build

GO := $(shell command -v go 2> /dev/null)
FS := /
# go source code files
GOSRCS=$(shell find . -name \*.go)
BUILDENV=CGO_ENABLED=0

ifeq ($(GO),)
  $(error could not find go. Is it in PATH? $(GO))
endif

ifneq ($(TARGET),)
  BUILDENV += GOOS=$(TARGET)
endif

tags: $(GOSRCS)
	gotags -R -f tags .

build: amocli

amocli: $(GOSRCS)
	@echo "--> Building amo console (amocli)"
	$(BUILDENV) go build ./cmd/amocli

install: $(GOSRCS)
	@echo "--> Installing amo console (amocli)"
	go install ./cmd/amocli

test:
	go test ./...

clean:
	rm -f amocli
