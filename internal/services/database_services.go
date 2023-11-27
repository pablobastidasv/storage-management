package services

import (
	"co.bastriguez/inventory/internal/models"
	"database/sql"
)

type TheService interface {
	InventoryService
	ProductService
}

type theService struct {
	db *sql.DB
}

func (t theService) RetrieveMainStorageContent() []models.InventoryContent {
	//TODO implement me
	panic("implement me")
}

func (t theService) RetrieveOpenRemissions() []models.Remission {
	//TODO implement me
	panic("implement me")
}

func (t theService) AddProductToMainStorage(productId string, qty int) error {
	//TODO implement me
	panic("implement me")
}

func (t theService) ListProducts() []models.Product {
	//TODO implement me
	panic("implement me")
}

func New(db *sql.DB) (TheService, error) {
	return &theService{
		db: db,
	}, nil
}
