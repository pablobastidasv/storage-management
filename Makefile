PHONY: run install generate build clean build/prod e2e/dev test


install:
	go install github.com/cosmtrek/air@latest
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest


run: 
	go run cmd/web-app/main.go


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

