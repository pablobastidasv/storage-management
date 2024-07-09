PHONY: run install generate build clean build/prod e2e/dev test generate


install:
	go install github.com/air-verse/air@latest
	go install github.com/vektra/mockery/v2@v2.43.2


run: 
	go run cmd/web-app/main.go


generate:
	@go generate ./...


build/prod:
	CGO_ENABLED=0 GOOS=linux go build -o dist/web-app cmd/web-app/main.go
	cp -r templates dist/.
	cp -r public/ dist/public


build: clean
	go build -o dist/web-app cmd/web-app/main.go
	cp -r templates dist/.
	cp -r public/ dist/public


clean:
	rm -rf dist


run/dev: run/db
	air


run/db:
	docker compose up -d


e2e/dev:
	cd e2e; npx playwright test --project chromium --ui


test: run/db
	go test ./...

