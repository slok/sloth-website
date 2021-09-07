
help: ## Show this help
	@echo "Help"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "    \033[36m%-20s\033[93m %s\n", $$1, $$2}'

.PHONY: default
default: help

.PHONY: run
run: ## Runs docs locally.
	./scripts/run.sh

.PHONY: gen
gen: sync  ## Generate site.
	./scripts/generate.sh

.PHONY: sync
sync: ## Sync all required files.
	./scripts/sync-sloth.sh
	./scripts/sync-sli-plugins.sh