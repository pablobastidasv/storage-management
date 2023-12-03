package repository

import (
	"co.bastriguez/inventory/internal/models"
	"database/sql"
)

type productRepo struct {
	db *sql.DB
}

func (p productRepo) FetchProducts() ([]models.Product, error) {
	rows, err := p.db.Query("select id, name, presentation  from products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		var presentation string
		err := rows.Scan(&product.Id, &product.Name, &presentation)
		if err != nil {
			return nil, err
		}
		product.Presentation = models.NewPresentation(presentation)

		products = append(products, product)
	}

	return products, nil
}

type ProductRepository interface {
	FetchProducts() ([]models.Product, error)
}

func NewProductsRepository(db *sql.DB) ProductRepository {
	return &productRepo{
		db,
	}
}
