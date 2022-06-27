.PHONY: build

build:
		go mod tidy
		go build -v ./cmd/main.go

migrate:
		migrate -path ./schema -database 'postgres://postgres:root@0.0.0.0:5432/postgres?sslmode=disable' up
.DEFAULT_GOAL:= build

