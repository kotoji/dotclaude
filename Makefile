.PHONY: help build install clean

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  %-10s %s\n", $$1, $$2}'

build: ## Build statusline binary
	go build -o bin/statusline ./cmd/statusline/

install: build ## Build and install (symlink to ~/.claude)
	bash setup.sh

clean: ## Remove built binaries
	rm -rf bin/
