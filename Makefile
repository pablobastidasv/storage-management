PHONY: run install generate build clean


install:
	cd public && pnpm install


generate:
	templ generate


run: generate
	go run cmd/web-app/main.go


build: clean
	go build -o dist/web-app cmd/web-app/main.go
	cp -r templates dist/.
	cd public && pnpm vite build
	mv public/dist dist/public


clean:
	rm -rf dist