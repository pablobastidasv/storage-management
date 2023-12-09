package repository

import (
	"co.bastriguez/inventory/internal/models"
	"context"
	"database/sql"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository interface {
	FetchProducts(ctx context.Context) ([]models.Product, error)
	ExistProductById(ctx context.Context, productId string) (bool, error)
}

// ------
type mongoProductRepo struct {
	collection *mongo.Collection
}

func (m mongoProductRepo) FetchProducts(ctx context.Context) ([]models.Product, error) {
	res, err := m.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var mongoProducts []product
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

func (m mongoProductRepo) ExistProductById(ctx context.Context, productId string) (bool, error) {
	find, err := m.collection.Find(ctx, bson.M{"_id": productId})
	if err != nil {
		return false, err
	}
	return find.Next(ctx), nil
}

// ------
type productRepo struct {
	db *sql.DB
}

func (p *productRepo) ExistProductById(ctx context.Context, productId string) (bool, error) {
	var exists sql.NullBool
	err := p.db.QueryRow("select exists(select 1 from products where id=$1);", productId).Scan(&exists)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func (p *productRepo) FetchProducts(ctx context.Context) ([]models.Product, error) {
	rows, err := p.db.Query("select id, name, presentation from products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		var presentation string
		err := rows.Scan(&product.Id, &product.Name, &presentation)
		if err != nil {
			return nil, err
		}
		product.Presentation = models.NewPresentation(presentation)

		products = append(products, product)
	}

	return products, nil
}

// --
func NewSqlProductsRepository(db *sql.DB) ProductRepository {
	return &productRepo{
		db,
	}
}

func NewMongoProductsRepository(client *mongo.Client) ProductRepository {
	// TODO: how to share the database name
	collection := client.Database("bastriguez").Collection("products")
	return &mongoProductRepo{
		collection: collection,
	}
}
