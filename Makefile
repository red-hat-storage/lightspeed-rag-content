DOCS_VERSION ?= 4.22
BINARY  := bin/odf-rag
IMAGE_CMD ?= podman
IMAGE_WITH_TAG ?= odf-ols-byok:latest

.PHONY: build download parse test clean all build-odf-rag

build:
	mkdir -p bin
	go build -buildvcs=false -o $(BINARY) ./cmd

download: build
	./$(BINARY) download --version $(DOCS_VERSION)

parse: build
	./$(BINARY) parse --version $(DOCS_VERSION)

test:
	go test ./...

clean:
	rm -rf bin/
	rm -rf parsed/

all: download parse

build-odf-rag: ## Builds a container with ODF RAG content
	$(IMAGE_CMD) build -f Containerfile \
		--build-arg DOC_VERSION=$(DOCS_VERSION) \
		-t $(IMAGE_WITH_TAG) .
