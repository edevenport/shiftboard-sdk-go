build:
	go build -o /dev/null

clean:
	go clean

test:
	go test

lint:
	golangci-lint run

.PHONY: build clean test clean
