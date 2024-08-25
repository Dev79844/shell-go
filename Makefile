build:
	@go build -o bin/sh

run:
	@./bin/sh

test:
	@go test ./.. -v
