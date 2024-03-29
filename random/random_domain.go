package random

import (
	"co.bastriguez/inventory/internal/models"
)

func WithProductId(id string) func(*models.InventoryProduct) {
	return func(i *models.InventoryProduct) {
		i.Id = id
	}
}

func InventoryProduct(options ...func(*models.InventoryProduct)) models.InventoryProduct {
	inv := models.InventoryProduct{
		Id:           Uuid(),
		Name:         String(),
		Presentation: models.Grms,
	}

	for _, o := range options {
		o(&inv)
	}

	return inv
}

func InventoryItem(options ...func(*models.InventoryItem)) models.InventoryItem {
	item := models.InventoryItem{
		Product: InventoryProduct(),
		Qty:     PositiveInt(),
	}

	for _, o := range options {
		o(&item)
	}

	return item
}
