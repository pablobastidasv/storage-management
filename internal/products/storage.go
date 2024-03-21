package products

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type ProductStorage interface {
	ListProducts(c context.Context) ([]Product, error)
}

type mongoProductStorage struct {
	db *mongo.Database
}

func NewMongoStorage(db *mongo.Database) ProductStorage {
	return &mongoProductStorage{
		db: db,
	}
}

// ListProducts implements ProductStorage.
func (*mongoProductStorage) ListProducts(c context.Context) ([]Product, error) {
	panic("unimplemented")
}
