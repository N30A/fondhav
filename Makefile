build:
	mkdir -p bin
	go build -o bin/ ./cmd/*

test:
	go test ./...

run: build
	./bin/fondhav

.PHONY: test build run
