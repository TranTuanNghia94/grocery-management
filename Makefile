# Makefile for grocery-management project

.PHONY: gen migrate build gen-dto

gen:
	go run ./tools/gen/main.go

migrate:
	go run ./tools/migration/main.go

build:
	go build -o ./bin/grocery-management ./cmd/grocery-management/main.go

gen-dto:
	openapi-generator-cli generate -g go-gin-server --additional-properties=packageName=dto   -o ./out -i /docs/swagger/openapi.yml

copy-dto:
	cp -r /docs/swagger/out/*.go ./internal/dto
