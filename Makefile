VERSION ?= 4.22
BINARY  := downloader

.PHONY: build download parse test clean all

build:
	go build -buildvcs=false -o $(BINARY) ./cmd

download: build
	./$(BINARY) download --version $(VERSION)

parse: build
	./$(BINARY) parse --version $(VERSION)

test:
	go test ./...

clean:
	rm -f $(BINARY)
	rm -rf parsed/

all: download parse
