package repository_test

import (
	"co.bastriguez/inventory/internal/models"
	"co.bastriguez/inventory/internal/repository"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"testing"
)

func createSomeProducts(ctx context.Context, collection *mongo.Collection, products []repository.Product) {
	_, err := collection.DeleteMany(ctx, bson.D{})
	if err != nil {
		log.Fatalf("error cleaning the collection: %s\n", err.Error())
	}

	var productsToInsert []interface{}
	for _, s := range products {
		productsToInsert = append(productsToInsert, s)
	}

	result, err := collection.InsertMany(ctx, productsToInsert)
	if err != nil {
		log.Fatalf("Error inserting many products %s\n", err.Error())
	}
	log.Printf("Ids: %s\n", result.InsertedIDs)
}

func randomProducts() []repository.Product {
	return []repository.Product{
		{
			Id:           "fe0b28ea-e96f-4f14-b0ea-4b7f6e0e6a59",
			Name:         "Copper",
			Presentation: models.Grms,
		},
		{
			Id:           "5cf1c718-a994-4673-aba4-b77bef39e7cd",
			Name:         "Bateries",
			Presentation: models.Grms,
		},
	}
}

func connect(ctx context.Context) (*mongo.Client, *mongo.Database) {
	// Making the connection
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:secretpassword@localhost:27017"))
	if err != nil {
		log.Fatalf("Error with the connection to mongo, %s\n", err.Error())
	}
	db := client.Database("bastriguez")

	return client, db
}

func createRandomProductWith(ctx context.Context, t *testing.T, collection *mongo.Collection, productId string) {
	prod := repository.Product{
		Id:           productId,
		Name:         "A name",
		Presentation: models.Grms,
	}

	createProduct(ctx, t, collection, &prod)
}

func createProduct(ctx context.Context, t *testing.T, collection *mongo.Collection, prod *repository.Product) {
	_, err := collection.InsertOne(ctx, prod)

	if err != nil {
		t.Fatalf("error creating a random product %s\n", err.Error())
	}
}

func createInventoryItem(ctx context.Context, t *testing.T, collection *mongo.Collection, m *models.InventoryItem) {
	inventoryProduct := repository.InventoryProduct{
		Id:           m.Product.Id,
		Name:         m.Product.Name,
		Presentation: m.Product.Presentation,
	}

	item := repository.InventoryItem{
		Product: inventoryProduct,
		Qty:     m.Qty,
	}

	_, err := collection.InsertOne(ctx, item)
	if err != nil {
		t.Fatalf("error inserting a just created inventory item %s\n", err.Error())
	}
}
