PHONY: run install generate build clean build/prod


install:
	go install github.com/cosmtrek/air@latest
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest


generate:
	templ generate


run: generate
	go run cmd/web-app/main.go


build/prod:
	CGO_ENABLED=0 GOOS=linux go build -o dist/web-app cmd/web-app/main.go
	cp -r templates dist/.
	cp -r public/ dist/public


migrate/prod:
	migrate -path database/migration/ -database $(DATABASE_URL) -verbose up


build: clean
	go build -o dist/web-app cmd/web-app/main.go
	cp -r templates dist/.
	cp -r public/ dist/public


clean:
	rm -rf dist


run/dev: run/db migrate/dev
	air


run/db:
	docker compose up -d


migrate/dev:
	migrate -path database/migration/ -database "postgresql://postgres:secretpassword@localhost:5432/?sslmode=disable" -verbose up