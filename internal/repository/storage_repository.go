package repository

import (
	"co.bastriguez/inventory/internal/models"
	"database/sql"
)

type StorageRepository interface {
	AddProduct(storage models.Storage, inventoryItem models.InventoryItem) error
	DecreaseProduct(storage models.Storage, inventoryItem models.InventoryItem) error
	CreateTransaction(transaction models.Transaction) error
	FetchItemsByStorage(storage *models.Storage) ([]models.InventoryItem, error)
}

type repository struct {
	db *sql.DB
}

func (r *repository) FetchItemsByStorage(_ *models.Storage) ([]models.InventoryItem, error) {
	rows, err := r.db.Query("select i.quantity, p.id, p.name, p.presentation from items i join public.products p on p.id = i.product_id")
	if err != nil {
		return nil, nil
	}
	defer rows.Close()

	var items []models.InventoryItem
	for rows.Next() {
		var item models.InventoryItem
		var presentation string
		if err := rows.Scan(&item.Qty, &item.Product.Id, &item.Product.Name, &presentation); err != nil {
			return nil, nil
		}
		item.Product.Presentation = models.NewPresentation(presentation)

		items = append(items, item)
	}

	return items, nil
}

func (r *repository) AddProduct(storage models.Storage, inventoryItem models.InventoryItem) error {
	//TODO implement me
	panic("implement me")
}

func (r *repository) DecreaseProduct(storage models.Storage, inventoryItem models.InventoryItem) error {
	//TODO implement me
	panic("implement me")
}

func (r *repository) CreateTransaction(transaction models.Transaction) error {
	//TODO implement me
	panic("implement me")
}

func NewStorageRepository(db *sql.DB) StorageRepository {
	return &repository{
		db: db,
	}
}
