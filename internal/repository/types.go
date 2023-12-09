package repository

import "co.bastriguez/inventory/internal/models"

type (
	Product struct {
		Id           string              `bson:"_id"`
		Name         string              `bson:"name"`
		Presentation models.Presentation `bson:"presentation"`
	}

	InventoryItem struct {
		Product InventoryProduct `bson:"product"`
		Qty     int              `bson:"qty"`
	}

	InventoryProduct struct {
		Id           string              `bson:"id"`
		Name         string              `bson:"name"`
		Presentation models.Presentation `bson:"presentation"`
	}
)
