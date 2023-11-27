package services

import "co.bastriguez/inventory/internal/models"

type ServiceError struct {
	code    string
	message string
}

type (
	InventoryService interface {
		RetrieveMainStorageContent() []models.InventoryContent
		RetrieveOpenRemissions() []models.Remission
		AddProductToMainStorage(productId string, qty int) error
	}

	ProductService interface {
		ListProducts() []models.Product
	}
)

func (s *ServiceError) Error() string {
	return s.message
}
