package repository

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"co.bastriguez/inventory/internal/models"
)

const ProductsCollectionName = "products"

type ProductRepository interface {
	PersistProduct(ctx context.Context, product *models.Product) error
	FetchProducts(ctx context.Context) ([]models.Product, error)
	ExistProductById(ctx context.Context, productId string) (bool, error)
	FindProduct(ctx context.Context, id string) (*models.Product, error)
}

// ------
type mongoProductRepo struct {
	collection *mongo.Collection
}

func (m *mongoProductRepo) PersistProduct(
	ctx context.Context,
	product *models.Product,
) error {
	toInsert := Product{
		Id:           product.Id.ToString(),
		Name:         product.Name.ToString(),
		Presentation: product.Presentation,
	}
	_, err := m.collection.InsertOne(ctx, toInsert)
	return err
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

	return models.NewProduct(
		prod.Id, prod.Name, prod.Presentation.ToString(),
	)
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
		prod, err := models.NewProduct(
			p.Id,
			p.Name,
			p.Presentation.ToString(),
		)
		if err != nil {
			return nil, err
		}
		products = append(products, *prod)
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
