build:
	@go build -o bin/footick main.go

run:
	@go run main.go

test:
	@go test -v -failfast ./...

clean:
	@rm -rf bin

.PHONY: run test clean
