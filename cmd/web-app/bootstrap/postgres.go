package bootstrap

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"

	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 8080

	dbUser = "postgres"
	dbPass = "secretpassword"
	dbHost = "localhost"
	dbPort = "5432"
	dbName = "bastriguez"
)

func dbOpen() (*sql.DB, error) {
	postgresURI := generatePostgresUriFromEnvvars()
	if env := os.Getenv(EnvVar); env == "" {
		slog.Debug("using local/dev database uri")
		postgresURI = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPass, dbHost, dbPort, dbName)
	}
	slog.Debug(fmt.Sprintf("URI string %s", postgresURI))
	return sql.Open("postgres", postgresURI)
}

func generatePostgresUriFromEnvvars() string {
	return os.Getenv("DB_URL")
}
