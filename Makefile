-include .env

SHELL            := /bin/sh
GOBIN            ?= $(GOPATH)/bin
PATH             := $(GOBIN):$(PATH)
GO               = go

M                = $(shell printf "\033[34;1m>>\033[0m")
TARGET_DIR       ?= $(PWD)/.build

ifeq ($(DELVE_ENABLED),true)
GCFLAGS	= -gcflags 'all=-N -l'
endif

.PHONY: install-tools
install-tools: $(GOBIN)
	@GOBIN=$(GOBIN) $(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@GOBIN=$(GOBIN) $(GO) install github.com/air-verse/air@latest

.PHONY: lint
lint:
	$(info $(M) running linters...)
	@$(GOBIN)/golangci-lint run -v --timeout 5m0s ./...

.PHONY: build
build:
	$(info $(M) building app...)
	@GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) build $(GCFLAGS) $(LDFLAGS) -o $(TARGET_DIR)/service ./cmd/main/*.go

.PHONY: start
start:
	go run cmd/main/main.go

.PHONY: watch
watch:
	$(info $(M) run...)
	@$(GOBIN)/air -c $(PWD)/.air.conf

%:
	@:
