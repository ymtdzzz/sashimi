IMAGE_NAME=bufbuild/buf
IMAGE_TAG=1.28.1
EXAMPLE_IMAGE_NAME=sashimi-example
EXAMPLE_IMAGE_TAG=latest
buf=docker run --volume "`pwd`:/workspace" --workdir /workspace $(IMAGE_NAME):$(IMAGE_TAG)

.PHONY: help
help:
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z0-9_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

buf-help: ## buf help
	$(buf) --help

buf-lint: ## buf lint proto files
	$(buf) lint -v

buf-fmt: ## buf format proto files
	$(buf) format

buf-gen-go: ## buf generate go proto files
	rm -rf ./gen/go
	$(buf) generate

buf-gen: ## buf generate proto files
	$(MAKE) buf-gen-go

build-example-image: ## build test docker image
	docker build -t $(EXAMPLE_IMAGE_NAME):$(EXAMPLE_IMAGE_TAG) -f Dockerfile.test .

run-example-job-a: ## run example job a
	docker run --rm $(EXAMPLE_IMAGE_NAME):$(EXAMPLE_IMAGE_TAG) ./job_runner.sh "./sashimi ./job_a_split 40"

run-example-job-b: ## run example job b
	docker run --rm $(EXAMPLE_IMAGE_NAME):$(EXAMPLE_IMAGE_TAG) ./job_runner.sh "./sashimi ./job_b_split 20"
