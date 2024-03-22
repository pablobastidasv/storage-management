package products

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductStorage interface {
	ListProducts(ctx context.Context) ([]Product, error)
}

type mongoProductStorage struct {
	collection *mongo.Collection
}

func NewMongoStorage(db *mongo.Database) ProductStorage {
	collection := db.Collection("products")
	return &mongoProductStorage{
		collection: collection,
	}
}

// ListProducts implements ProductStorage.
func (m *mongoProductStorage) ListProducts(ctx context.Context) ([]Product, error) {
	res, err := m.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var products []Product
	if err := res.All(ctx, &products); err != nil {
		return nil, err
	}

	return products, nil

}
