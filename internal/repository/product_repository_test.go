package repository_test

import (
	"co.bastriguez/inventory/internal/models"
	"co.bastriguez/inventory/internal/repository"
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"testing"
	"time"
)

func TestMongoProductRepo_FetchProducts(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, database := connect(ctx)
	collection := database.Collection(repository.ProductsCollectionName)

	givenProducts := randomProducts()
	persistProducts(ctx, collection, givenProducts)

	var expected []models.Product
	for _, prod := range givenProducts {
		expected = append(expected, models.Product{
			Id:           prod.Id,
			Name:         prod.Name,
			Presentation: prod.Presentation,
		})
	}

	sut := repository.NewMongoProductsRepository(database)

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
	collection := database.Collection(repository.ProductsCollectionName)

	existingProductId := randomProductId(t)
	_ = createRandomProductWith(ctx, t, collection, existingProductId)

	nonExistingProductId := randomProductId(t)

	type fields struct {
		database *mongo.Database
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
			fields: fields{database: database},
			args: args{
				ctx:       ctx,
				productId: existingProductId,
			},
			want:    true,
			wantErr: assert.NoError,
		},
		{
			name:   "Given a NON exising product, then return false",
			fields: fields{database: database},
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
			sut := repository.NewMongoProductsRepository(tt.fields.database)
			got, err := sut.ExistProductById(tt.args.ctx, tt.args.productId)
			if !tt.wantErr(t, err, fmt.Sprintf("ExistProductById(%v, %v)", tt.args.ctx, tt.args.productId)) {
				return
			}
			assert.Equalf(t, tt.want, got, "ExistProductById(%v, %v)", tt.args.ctx, tt.args.productId)
		})
	}
}

func Test_mongoProductRepo_FindProduct(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	_, database := connect(ctx)
	collection := database.Collection(repository.ProductsCollectionName)

	existingProductId := randomProductId(t)
	persistedProduct := createRandomProductWith(ctx, t, collection, existingProductId)
	expectedProduct := models.Product{
		Id:           persistedProduct.Id,
		Name:         persistedProduct.Name,
		Presentation: persistedProduct.Presentation,
	}

	nonExistingProductId := randomProductId(t)

	type fields struct {
		collection *mongo.Collection
	}
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Product
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:   "when product exists, return the product",
			fields: fields{collection: collection},
			args: args{
				ctx: ctx,
				id:  existingProductId,
			},
			want:    &expectedProduct,
			wantErr: assert.NoError,
		},
		{
			name:   "when product does not exists, return a nil value",
			fields: fields{collection: collection},
			args: args{
				ctx: ctx,
				id:  nonExistingProductId,
			},
			want:    nil,
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sut := repository.NewMongoProductsRepository(database)
			got, err := sut.FindProduct(tt.args.ctx, tt.args.id)
			if !tt.wantErr(t, err, fmt.Sprintf("FindProduct(%v, %v)", tt.args.ctx, tt.args.id)) {
				return
			}
			assert.Equalf(t, tt.want, got, "FindProduct(%v, %v)", tt.args.ctx, tt.args.id)
		})
	}
}
