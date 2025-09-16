.ONESHELL:
.DELETE_ON_ERROR:
.SHELLFLAGS := -eu -o pipefail -c
SHELL := bash
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules
MAKEFLAGS += --no-print-directory

DOC_SLUG ?= "red_hat_openshift_data_foundation"
DOC_VERSION ?= "4.19"
VECTOR_DB ?= "vector_db"
INDEX ?= "odf-${DOC_VERSION}"
IMAGE_CMD ?= "podman"
IMAGE_WITH_TAG ?= "odf-ols-byok:latest"

##@ Commands
.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {
		FS = ":.*##";
		printf "\nUsage:\n  make \033[36m<target>\033[0m\n"
	}
	/^[a-zA-Z_0-9-]+:.*?##/ {
		printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2
	}
	/^##@/ {
		printf "\n\033[1m%s\033[0m\n", substr($$0, 5)
	}' $(MAKEFILE_LIST)

.PHONY: setup-env
setup-env: ## Installs python environment and required python packages into a virtual env
	@uv python install 3.11
	@uv venv --python 3.11
	@uv pip install -r ../requirements.cpu.txt

.PHONY: download-docs
download-docs: ## Downloads ODF docs into 'odf' folder corresponding to the version and preps them
	@pushd ..

	@mkdir -p odf/${DOC_SLUG}/${DOC_VERSION}/downloads/

	@python3.11 scripts/doc_downloader/downloader.py \
	--doc-url "https://docs.redhat.com/en/documentation/${DOC_SLUG}/${DOC_VERSION}/" \
	--output-dir odf/${DOC_SLUG}/${DOC_VERSION}/downloads/ \
	--max-retries 1

	@find odf/${DOC_SLUG}/${DOC_VERSION}/downloads/html-single -empty -type d -delete
	@find odf/${DOC_SLUG}/${DOC_VERSION}/downloads/html-single -not -name index -type f -delete

	@find odf/${DOC_SLUG}/${DOC_VERSION}/downloads/html -type f -delete
	@find odf/${DOC_SLUG}/${DOC_VERSION}/downloads/html -empty -type d -delete

	@find odf/${DOC_SLUG}/${DOC_VERSION}/downloads/html-single -name index -type f -exec mv {} {}.html \;

	@mv odf/${DOC_SLUG}/${DOC_VERSION}/downloads/html-single/* odf/${DOC_SLUG}/${DOC_VERSION}/downloads/

	@popd

.PHONY: create-rag
create-rag: ## Creates vector db based on cached downloads
	@pushd ..
	@python3.11 scripts/html_embeddings/generate_embeddings.py \
	--doc-url-version ${DOC_VERSION} \
	--doc-url-slug ${DOC_SLUG} \
	--output-dir ${VECTOR_DB} \
	--index ${INDEX} \
	--skip-runbooks \
	--use-cached-downloads \
	--cache-dir odf/ \
	--verbose
	@popd

.PHONY: build-odf-rag
build-odf-rag: ## Builds a container with ODF RAG content
	@pushd ..
	@${IMAGE_CMD} build -f odf/Containerfile -t ${IMAGE_WITH_TAG} .
	@popd
