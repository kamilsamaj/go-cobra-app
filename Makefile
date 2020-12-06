XC_OS="linux darwin windows"
#XC_ARCH="amd64 arm64" # darwin amr64 support coming in go 1.16
XC_ARCH="amd64"
XC_PARALLEL="4"
BIN="./bin"
SRC=$(shell find . -name "*.go")

ifeq (, $(shell which gox))
$(warning "could not find gox in $(PATH), run: go get github.com/mitchellh/gox")
endif

.PHONY: all build

default: all

all: build

build:
	CGO_ENABLED=0 gox \
		-os=$(XC_OS) \
		-arch=$(XC_ARCH) \
		-parallel=$(XC_PARALLEL) \
		-output=$(BIN)/{{.Dir}}-{{.OS}}-{{.Arch}} \
		;
