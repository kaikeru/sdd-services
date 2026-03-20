.PHONY: build test lint run-orchestrator

build:
	go build ./...

test:
	go test ./...

lint:
	go vet ./...

run-orchestrator:
	go run ./cmd/kimitsu start orchestrator
