package databases

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

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
