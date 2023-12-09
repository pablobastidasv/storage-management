package repository_test

import (
	"co.bastriguez/inventory/internal/models"
	"co.bastriguez/inventory/internal/repository"
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"testing"
	"time"
)

func TestMongoProductRepo_FetchProducts(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, database := connect(ctx)
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
	collection := database.Collection("products")

	existingProductId := randomProductId(t)
	createRandomProductWith(ctx, t, collection, existingProductId)

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

func randomProductId(t *testing.T) string {
	existingProductId, err := uuid.NewUUID()
	if err != nil {
		t.Fatalf("error generating the uuid %s\n", err.Error())
	}
	return existingProductId.String()
}
