.PHONY: build lint test integration-test run-life

build:
	go build ./...

lint:
	exit `go fmt ./... | wc -l`

test:
	go test ./...

integration-test:
	go run ./cmd/integration/

run-life:
	go run ./cmd/life/
