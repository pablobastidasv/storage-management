package repository

import (
	"co.bastriguez/inventory/internal/models"
	"database/sql"
	"errors"
)

type productRepo struct {
	db *sql.DB
}

func (p *productRepo) ExistProductById(productId string) (bool, error) {
	var exists sql.NullBool
	err := p.db.QueryRow("select exists(select 1 from products where id=$1);", productId).Scan(&exists)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func (p *productRepo) FetchProducts() ([]models.Product, error) {
	rows, err := p.db.Query("select id, name, presentation from products")
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
	ExistProductById(productId string) (bool, error)
}

func NewProductsRepository(db *sql.DB) ProductRepository {
	return &productRepo{
		db,
	}
}
