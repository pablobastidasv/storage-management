PHONY: run install generate build clean build/prod e2e/dev test generate live


install:
	go install github.com/air-verse/air@latest
	go install github.com/vektra/mockery/v2@v2.43.2
	go install github.com/a-h/templ/cmd/templ@latest
	pnpm install -D tailwindcss
	brew install golang-migrate 


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


live/templ:
	templ generate --watch --proxy="http://localhost:8080" --open-browser=false -v


live/server:
	air


live: run/db migrate/local
	make -j5 live/templ live/server


run/db:
	docker compose up -d


e2e/dev:
	cd e2e; npx playwright test --project chromium --ui


test: run/db
	go test ./...


migrate: 
	migrate -database ${POSTGRESQL_URL} -path db/migrations up


migrate/local:
	make migrate POSTGRESQL_URL=postgres://postgres:secretpassword@localhost:5432/bastriguez?sslmode=disable
