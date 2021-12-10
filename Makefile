MAKEFLAGS += --warn-undefined-variables
SHELL := /bin/bash
ifeq ($(word 1,$(subst ., ,$(MAKE_VERSION))),4)
.SHELLFLAGS := -eu -o pipefail -c
endif
.ONESHELL:

.DEFAULT_GOAL := test
PACKAGE := $(shell go list)

.PHONY: help
help:  ## Show this help
	@grep '.*:.*##' Makefile | grep -v grep  | sort | sed 's/:.* ##/:/g' | column -t -s:

.PHONY: run
run:
	go run main.go

.PHONY: test
test:  ## Run go test
	go test -v ./...
