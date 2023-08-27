include .env
export

dev:
	go run main.go

start:
	go build main.go && ./main