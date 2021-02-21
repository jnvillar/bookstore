build:
	go build -o bookstore main.go

run:
	./bookstore

lint:
	go run github.com/golangci/golangci-lint/cmd/golangci-lint run