package repository

import (
	"co.bastriguez/inventory/internal/models"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const ProductsCollectionName = "products"

type ProductRepository interface {
	PersistProduct(ctx context.Context) (*models.Product, error)
	FetchProducts(ctx context.Context) ([]models.Product, error)
	ExistProductById(ctx context.Context, productId string) (bool, error)
	FindProduct(ctx context.Context, id string) (*models.Product, error)
}

// ------
type mongoProductRepo struct {
	collection *mongo.Collection
}

func (m *mongoProductRepo) PersistProduct(ctx context.Context) (*models.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (m *mongoProductRepo) FindProduct(ctx context.Context, id string) (*models.Product, error) {
	var prod Product
	err := m.collection.FindOne(ctx, bson.D{{"_id", id}}).Decode(&prod)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}

	product := toProdModel(&prod)
	return product, nil
}

func toProdModel(p *Product) *models.Product {
	return &models.Product{
		Id:           p.Id,
		Name:         p.Name,
		Presentation: p.Presentation,
	}
}

func (m *mongoProductRepo) FetchProducts(ctx context.Context) ([]models.Product, error) {
	res, err := m.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var mongoProducts []Product
	if err := res.All(ctx, &mongoProducts); err != nil {
		return nil, err
	}

	var products []models.Product
	for _, p := range mongoProducts {
		products = append(products, models.Product{
			Id:           p.Id,
			Name:         p.Name,
			Presentation: p.Presentation,
		})
	}
	return products, nil
}

func (m *mongoProductRepo) ExistProductById(ctx context.Context, productId string) (bool, error) {
	find, err := m.collection.Find(ctx, bson.M{"_id": productId})
	if err != nil {
		return false, err
	}
	return find.Next(ctx), nil
}

func NewMongoProductsRepository(database *mongo.Database) ProductRepository {
	collection := database.Collection("products")
	return &mongoProductRepo{
		collection: collection,
	}
}
