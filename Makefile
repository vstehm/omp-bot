.PHONY: run
run:
	go run cmd/bot/main.go

.PHONY: test
test:
	go test -v ./...

.PHONY: build
build:
	go build -o bot cmd/bot/main.go

.PHONY: lint
lint:
	golangci-lint run ./...