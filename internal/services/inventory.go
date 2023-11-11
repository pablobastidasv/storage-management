package services

import "co.bastriguez/inventory/internal/models"

type (
	InventoryService interface {
		RetrieveInventory() []models.InventoryContent
		RetrieveOpenRemissions() []models.Remission
	}
)
