package databases

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func NewRelational() (*sql.DB, error) {
	return newConnection()
}

func NewMongo() (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	uri := "mongodb://root:secretpassword@localhost:27017"

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("Error with the connection to mongo, %s\n", err.Error())
	}

	return client.Database("bastriguez"), nil
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
