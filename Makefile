.PHONY: build lint test integrate

build:
	go build ./...

lint:
	exit `go fmt ./... | wc -l`

test:
	go test ./...

integrate:
	go run ./pkg/integrate/
