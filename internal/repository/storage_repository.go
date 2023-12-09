package repository

import (
	"co.bastriguez/inventory/internal/models"
	"context"
	"database/sql"
	"errors"
)

type StorageRepository interface {
	FetchItemsByStorage(ctx context.Context, storage *models.Storage) ([]models.InventoryItem, error)
	FindItemBy(storageId string, productId string) (*models.InventoryItem, error)
	UpdateItem(storageId string, item *models.InventoryItem) error
	FindMainStorage() (*models.Storage, error)
}

type repository struct {
	db *sql.DB
}

func (r *repository) FindMainStorage() (*models.Storage, error) {
	row := r.db.QueryRow("select s.id from storages s limit 1")
	var storage models.Storage
	err := row.Scan(&storage.Id)
	if err != nil {
		return nil, err
	}

	return &storage, nil
}

func (r *repository) FindItemBy(storageId string, productId string) (*models.InventoryItem, error) {
	row := r.db.QueryRow(
		`select i.quantity, p.name, p.presentation from items i 
			join public.products p on p.id = i.product_id 
			where i.storage_id = $1 
			and i.product_id = $2`,
		storageId,
		productId,
	)

	var item models.InventoryItem
	var presentation string
	err := row.Scan(&item.Qty, &item.Product.Name, &presentation)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	item.Product.Id = productId
	item.Product.Presentation = models.NewPresentation(presentation)

	return &item, nil
}

func (r *repository) UpdateItem(storageId string, item *models.InventoryItem) error {
	_, err := r.db.Exec(`
		INSERT INTO items(storage_id, product_id, quantity)
			VALUES($1, $2, $3) 
			ON CONFLICT (storage_id, product_id) 
			DO 
			   UPDATE SET quantity = $3;
		`, storageId, item.Product.Id, item.Qty)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) FetchItemsByStorage(_ context.Context, _ *models.Storage) ([]models.InventoryItem, error) {
	rows, err := r.db.Query(`select i.quantity, p.id, p.name, p.presentation 
								 from items i 
								 join public.products p on p.id = i.product_id
								 order by p.name`)
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

func NewStorageRepository(db *sql.DB) StorageRepository {
	return &repository{
		db: db,
	}
}
