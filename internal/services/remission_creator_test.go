package services_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/mock"

	"co.bastriguez/inventory/internal/models"
	"co.bastriguez/inventory/internal/services"
	"co.bastriguez/inventory/random"
)

type (
	RemissionCreatorDataPortMock struct {
		mock.Mock
	}
)

func (r *RemissionCreatorDataPortMock) FindInventoryItem(
	ctx context.Context,
	storageId models.StorageId,
	productId models.ProductId,
) (*models.InventoryItem, error) {
	args := r.Called(ctx, storageId, productId)
	return args.Get(0).(*models.InventoryItem), args.Error(1)
}

func (r *RemissionCreatorDataPortMock) UpdateInventoryItem(
	ctx context.Context,
	inventoryItem *models.InventoryItem,
) error {
	args := r.Called(ctx, inventoryItem)
	return args.Error(0)
}

func Test_CreateRemission_inventoryItemAmmountOfProductDecrease(t *testing.T) {
	mock := new(RemissionCreatorDataPortMock)
	sut := services.NewRemissionCreator(mock)

	ctx := context.Background()
	storageId := random.Uuid()
	productId := random.Uuid()
	qty := 6

	initialQty := func(i *models.InventoryItem) {
		i.Qty = 10
	}
	inventoryItem := random.InventoryItem(initialQty)

	expectedInventoryItem := inventoryItem
	expectedInventoryItem.Qty = 4

	mock.
		On(
			"FindInventoryItem",
			ctx,
			models.StorageId(storageId),
			models.ProductId(productId),
		).Return(&inventoryItem, nil).Once().
		On(
			"UpdateInventoryItem",
			ctx,
			&expectedInventoryItem,
		).Return(nil).Once()

	input := services.NewRemissionCreatorInput(storageId, productId, qty)

	sut.CreateRemission(ctx, input)

	mock.AssertExpectations(t)
}
