PHONY: run install generate build clean


install:
	cd public && pnpm install
	go install github.com/cosmtrek/air@latest


generate:
	templ generate


run: generate
	go run cmd/web-app/main.go


build: clean
	go build -o dist/web-app cmd/web-app/main.go
	cp -r templates dist/.
	cp -r public/ dist/public


clean:
	rm -rf dist


run/dev:
	air