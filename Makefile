.PHONY: run
run: build ## run latencyd, you can specify the commandline args via env var ARGS
	@build/latencyd

.PHONY: build
build: ## Build the latencyd binary
	@go build -o build/latencyd cmd/latencyd/latencyd.go

.PHONY: clean
clean: ## Remove build artifacts
	@rm -rf build/*

.PHONY: help
help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

