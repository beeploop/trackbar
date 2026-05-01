build:
	@go build -o bin/trackbar main.go

run:
	@go run main.go

test:
	@grc go test -v -failfast ./...

clean:
	@rm -rf bin

.PHONY: run test clean
