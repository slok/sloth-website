VERSION = latest
IMAGE_NAME ?=  slok.dev/sloth-website

BUILD_CMD := IMAGE=${IMAGE_NAME} VERSION=${VERSION} DOCKER_FILE_PATH=./Dockerfile ./scripts/build.sh
RUN_ON_DOCKER := docker run --rm -it -p 8000:8000 -v "${PWD}":/docs --entrypoint="" ${IMAGE_NAME}

help: ## Show this help
	@echo "Help"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "    \033[36m%-20s\033[93m %s\n", $$1, $$2}'

.PHONY: default
default: help

.PHONY: run
run: ## Runs docs locally.
	$(RUN_ON_DOCKER) /bin/sh -c './scripts/run.sh'

.PHONY: build
build:  ## Build docker image.
	$(BUILD_CMD)

.PHONY: gen
gen: ## Generate site.
	$(RUN_ON_DOCKER) /bin/sh -c './scripts/generate.sh'

.PHONY: sync
sync: ## Sync all required files.
	./scripts/sync-sloth.sh
	./scripts/sync-sli-plugins.sh