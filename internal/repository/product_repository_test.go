package repository

import (
	"co.bastriguez/inventory/internal/models"
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"testing"
	"time"
)

func TestMongoProductRepo_FetchProducts(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	client, database := connect(ctx)
	collection := database.Collection("products")

	givenProducts := randomProducts()
	createSomeProducts(ctx, collection, givenProducts)

	var expected []models.Product
	for _, prod := range givenProducts {
		expected = append(expected, models.Product{
			Id:           prod.Id,
			Name:         prod.Name,
			Presentation: prod.Presentation,
		})
	}

	sut := NewMongoProductsRepository(client)

	result, err := sut.FetchProducts(ctx)
	if err != nil {
		t.Fatalf("error fetching products %s\n:", err.Error())
	}

	assert.Equal(t, expected, result)
}

func Test_mongoProductRepo_ExistProductById(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, database := connect(ctx)
	collection := database.Collection("products")

	existingProductId := randomProductId(t)
	createRandomProductWith(ctx, t, collection, existingProductId)

	nonExistingProductId := randomProductId(t)

	type fields struct {
		collection *mongo.Collection
	}
	type args struct {
		ctx       context.Context
		productId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:   "Given a exising product, then return true",
			fields: fields{collection: collection},
			args: args{
				ctx:       ctx,
				productId: existingProductId,
			},
			want:    true,
			wantErr: assert.NoError,
		},
		{
			name:   "Given a NON exising product, then return false",
			fields: fields{collection: collection},
			args: args{
				ctx:       ctx,
				productId: nonExistingProductId,
			},
			want:    false,
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := mongoProductRepo{
				collection: tt.fields.collection,
			}
			got, err := m.ExistProductById(tt.args.ctx, tt.args.productId)
			if !tt.wantErr(t, err, fmt.Sprintf("ExistProductById(%v, %v)", tt.args.ctx, tt.args.productId)) {
				return
			}
			assert.Equalf(t, tt.want, got, "ExistProductById(%v, %v)", tt.args.ctx, tt.args.productId)
		})
	}
}

func randomProductId(t *testing.T) string {
	existingProductId, err := uuid.NewUUID()
	if err != nil {
		t.Fatalf("error generating the uuid %s\n", err.Error())
	}
	return existingProductId.String()
}

// ==== Utility functions
func createSomeProducts(ctx context.Context, collection *mongo.Collection, products []product) {
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

func randomProducts() []product {
	return []product{
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
	_, err := collection.InsertOne(ctx, product{
		Id:           productId,
		Name:         "A name",
		Presentation: models.Grms,
	})

	if err != nil {
		t.Fatalf("error creating a random product %s\n", err.Error())
	}
}
