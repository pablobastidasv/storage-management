package repository

import (
	"co.bastriguez/inventory/internal/models"
	"database/sql"
)

type StorageRepository interface {
	AddProduct(storage models.Storage, inventoryItem models.InventoryItem) error
	DecreaseProduct(storage models.Storage, inventoryItem models.InventoryItem) error
	CreateTransaction(transaction models.Transaction) error
}

type repository struct {
	db *sql.DB
}

func (r repository) AddProduct(storage models.Storage, inventoryItem models.InventoryItem) error {
	//TODO implement me
	panic("implement me")
}

func (r repository) DecreaseProduct(storage models.Storage, inventoryItem models.InventoryItem) error {
	//TODO implement me
	panic("implement me")
}

func (r repository) CreateTransaction(transaction models.Transaction) error {
	//TODO implement me
	panic("implement me")
}

func New(db *sql.DB) StorageRepository {
	return &repository{
		db: db,
	}
}
