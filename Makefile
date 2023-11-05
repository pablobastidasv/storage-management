PHONY: run install generate

install:
	go install github.com/a-h/templ/cmd/templ@latest

generate:
	templ generate

run: generate
	go run cmd/web-app/main.go