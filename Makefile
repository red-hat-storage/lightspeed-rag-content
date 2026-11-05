VERSION ?= 4.22
BINARY  := bin/odf-rag
IMAGE_CMD ?= podman
IMAGE_WITH_TAG ?= odf-ols-byok:latest

.PHONY: build download parse test clean all build-odf-rag

build:
	mkdir -p bin
	go build -buildvcs=false -o $(BINARY) ./cmd

download: build
	./$(BINARY) download --version $(VERSION)

parse: build
	./$(BINARY) parse --version $(VERSION)

test:
	go test ./...

clean:
	rm -rf bin/
	rm -rf parsed/

all: download parse

build-odf-rag: ## Builds a container with ODF RAG content
	$(IMAGE_CMD) build -f Containerfile \
		--build-arg DOC_VERSION=$(VERSION) \
		--build-arg VECTOR_DB_INDEX=odf_$(subst .,_,$(VERSION)) \
		-t $(IMAGE_WITH_TAG) .
