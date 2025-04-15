# Makefile for grocery-management project

.PHONY: gen build

gen:
	go run ./tools/gen/main.go

build:
	go build -o ./bin/grocery-management ./cmd/grocery-management/main.go


