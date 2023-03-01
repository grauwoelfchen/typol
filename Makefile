GOTEST := $(shell if type gotest > /dev/null 2>&1; then echo "gotest"; \
	else echo "go test"; fi)

check:
	@go vet ./...
.PHONY: check

fmt:
	@out=`gofmt -l . 2>&1`; \
		if [ "$$out" ]; then \
			echo "Run \`gofmt\` for the followings:"; \
			echo "$$out"; \
			exit 1; \
		fi
.PHONY: fmt

lint:
	@golangci-lint run --out-format=line-number ./...
.PHONY: lint

test:
	@$(GOTEST) -v ./typol/...
.PHONY: test

test\:integration: build
	@$(GOTEST) -v -tags integration ./cmd/...
.PHONY: test\:integration

build:
	@go build -o ./dst/ ./...
.PHONY: build

clean:
	@go clean
	@rm -f ./dst/*
.PHONY: clean
