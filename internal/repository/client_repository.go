package repository

import (
	"context"

	"co.bastriguez/inventory/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type ClientRepository interface {
	SaveClient(context.Context, models.Client) error
}

const ClientsCollectionName = "clients"

type mongoClientRepository struct {
	collection *mongo.Collection
}

func NewClientMongoRepository(db *mongo.Database) ClientRepository {
	collection := db.Collection(ClientsCollectionName)

	return &mongoClientRepository{
		collection: collection,
	}
}

func (m mongoClientRepository) SaveClient(ctx context.Context, client models.Client) error {
	_, err := m.collection.InsertOne(ctx, client)

	return err
}
