package services

import "co.bastriguez/inventory/internal/entities"

type (
	InventoryService interface {
		RetrieveInventory() []entities.InventoryContent
	}
)
