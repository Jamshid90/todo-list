-include .env
export

CURRENT_DIR=$(shell pwd)
CMD_DIR=./cmd

.PHONY: run
run:
	go run cmd/main.go

# build for current os
.PHONY: build
build:
	CGO_ENABLED=0 GOARCH="amd64" GOOS=linux go build -ldflags="-s -w" -o ./bin/server ${CMD_DIR}/main.go

# generate swagger	
.PHONY: swagger-gen
swagger-gen:
	swag init --dir ./internal/delivery/http -g router.go -o ./internal/delivery/http/docs

# migrate
.PHONY: migrate
migrate:
	migrate -source file://migrations -database postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DATABASE}?sslmode=disable up

.PHONY: start
start:
	docker-compose up -d --build

.PHONY: stop
stop:
	docker-compose down	
