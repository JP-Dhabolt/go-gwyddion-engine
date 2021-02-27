.PHONY: build lint test integration-test

build:
	go build ./...

lint:
	exit `go fmt ./... | wc -l`

test:
	go test ./...

integration-test:
	go run ./cmd/integration/
