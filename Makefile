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

build:
	@go build -o ./dst/ ./...
.PHONY: build

clean:
	@go clean
.PHONY: clean
