build:
	@go build -o bin/api

run:
	make build
	@./bin/api

test:
	@go test -v ./...

test-watch:
	@gotestsum --watch --format testname --hide-summary=all

seed:
	go run scripts/seed.go
