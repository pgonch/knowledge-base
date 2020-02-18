GOENV = CGO_ENABLED=0 GO111MODULE=on GOPRIVATE=github.com/pgonch
GO = $(GOENV) go

all: mod tools-all gen lint build
.PHONY: all

build:
	$(GO) build -o ./bin/knowledge-base    .
.PHONY: build

mod:
	$(GO) mod download
mod-tidy:
	$(GO) mod tidy
.PHONY: mod mod-tidy

tools-go-raml:
	GO111MODULE=off go get github.com/Jumpscale/go-raml

tools-test:
	$(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint

tools-all: tools-go-raml tools-test
.PHONY: tools-go-raml tools-test tools-all

gen:
	$(GO) generate .
.PHONY: gen

lint:
	$(GOENV) golangci-lint run .
.PHONY: lint

clean:
	$(GO) clean
	rm -rf bin static dist goraml types routes.go \
	*_if.go
.PHONY: clean
