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

func TestMongoRepository_FindItemByProductId(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, database := connect(ctx)
	collection := database.Collection(repository.InventoryItemsCollectionName)

	existingProductId := randomProductId(t)
	expectedItem := models.InventoryItem{
		Product: models.InventoryProduct{
			Id:           existingProductId,
			Name:         "The Pavo",
			Presentation: models.Amount,
		},
		Qty: 42,
	}
	createInventoryItem(ctx, t, collection, &expectedItem)
	nonExistingProductId := randomProductId(t)

	type fields struct {
		database *mongo.Database
	}
	type args struct {
		ctx       context.Context
		in1       string
		productId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.InventoryItem
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:   "when product has inventory items, then return the item in the inventory of that product",
			fields: fields{database: database},
			args: args{
				ctx:       ctx,
				in1:       "",
				productId: existingProductId,
			},
			want:    &expectedItem,
			wantErr: assert.NoError,
		},
		{
			name:   "when product does not have items in the inventory, then return nil",
			fields: fields{database: database},
			args: args{
				ctx:       ctx,
				in1:       "",
				productId: nonExistingProductId,
			},
			want:    nil,
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sut := repository.NewStorageMongoRepository(tt.fields.database)
			got, err := sut.FindItemByProductId(tt.args.ctx, tt.args.in1, tt.args.productId)
			if !tt.wantErr(t, err, fmt.Sprintf("FindItemByProductId(%v, %v, %v)", tt.args.ctx, tt.args.in1, tt.args.productId)) {
				return
			}
			assert.Equalf(t, tt.want, got, "FindItemByProductId(%v, %v, %v)", tt.args.ctx, tt.args.in1, tt.args.productId)
		})
	}
}

func TestMongoRepository_FetchItemsByStorage(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, database := connect(ctx)
	collection := database.Collection(repository.InventoryItemsCollectionName)

	expectedItems := randomInventoryItems(t)
	cleanCollection(ctx, t, collection)
	createRandomInventoryItemsWith(ctx, t, collection, expectedItems)

	sut := repository.NewStorageMongoRepository(database)

	result, err := sut.FetchItemsByStorage(ctx, nil)
	if err != nil {
		t.Fatalf("error fetching the items %s\n", err.Error())
	}

	assert.Equal(t, expectedItems, result)
}
