package services

import (
	"context"

	"co.bastriguez/inventory/internal/models"
)

type (
	RemissionCreator struct {
		dataPort RemissionCreatorDataPort
	}

	RemissionCreatorInput struct {
		storageId string
		productId string
		qty       int
	}

	RemissionCreatorDataPort interface {
		FindInventoryItem(
			context.Context,
			models.StorageId,
			models.ProductId,
		) (*models.InventoryItem, error)
		UpdateInventoryItem(context.Context, *models.InventoryItem) error
	}
)

func NewRemissionCreator(dataPort RemissionCreatorDataPort) *RemissionCreator {
	return &RemissionCreator{
		dataPort: dataPort,
	}
}

func NewRemissionCreatorInput(storageId string, productId string, qty int) RemissionCreatorInput {
	return RemissionCreatorInput{
		storageId: storageId,
		productId: productId,
		qty:       qty,
	}
}

func (r *RemissionCreator) CreateRemission(
	ctx context.Context,
	input RemissionCreatorInput,
) {
	storageId, _ := models.StorageIdFrom(input.storageId)
	productId, _ := models.ProductIdFrom(input.productId)

	inventoryItem, _ := r.dataPort.FindInventoryItem(ctx, storageId, productId)

	inventoryItem.Qty -= input.qty

	r.dataPort.UpdateInventoryItem(ctx, inventoryItem)
}
