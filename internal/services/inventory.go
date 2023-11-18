package services

import "co.bastriguez/inventory/internal/models"

type (
	InventoryService interface {
		RetrieveProducts() []models.InventoryContent
		RetrieveOpenRemissions() []models.Remission
	}
)
