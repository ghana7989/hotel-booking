build:
	@go build -o bin/api

run:
	make build
	@./bin/api

test:
	@go test -v ./...
