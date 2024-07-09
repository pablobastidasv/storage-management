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

	dbUser = "inventory"
	dbPass = "inventory"
	dbHost = "localhost"
	dbPort = "5432"
	dbName = "inventory"
)

func dbOpen() (*sql.DB, error) {
	postgresURI := generatePostgresUriFromEnvvars()
	if env := os.Getenv(EnvVar); env != "" {
        slog.Debug("using local/dev database uri")
		postgresURI = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPass, dbHost, dbPort, dbName)
	}
	return sql.Open("postgres", postgresURI)
}

func generatePostgresUriFromEnvvars() string {
	return os.Getenv("DB_URL")
}
