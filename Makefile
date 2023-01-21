GO               = go
GOBIN			 ?= $(PWD)/bin
TARGET_DIR       ?= $(PWD)/.build

.PHONY: all
all: run

.PHONY: run
run: ## run 'server' binary
	$(info $(M) run server...)
	GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) run ./cmd/server/*.go

