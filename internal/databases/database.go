package databases

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func New() (*sql.DB, error) {
	return newConnection()
}

func newConnection() (*sql.DB, error) {
	host := "localhost"
	port := 5432
	user := "postgres"
	password := "secretpassword"
	dbname := "postgres"

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("error opening the connection, %s", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("error when pinging to the database, %s", err)
	}

	return db, nil
}
