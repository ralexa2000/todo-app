.PHONY: run
run:
	go run cmd/main/main.go

build:
	go build -o bin/bot cmd/main/main.go
