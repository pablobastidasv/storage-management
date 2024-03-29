package services

import (
	"context"

	"co.bastriguez/inventory/internal/models"
	"co.bastriguez/inventory/internal/repository"
)

type (
	productService struct {
		productRepo repository.ProductRepository
	}

	ProductList struct {
		Items []models.Product
	}

	ProductOverview struct {
		Id           string
		Name         string
		Presentation models.Presentation
	}
)

// DeleteProduct implements ProductsService.
func (*productService) DeleteProduct(ctx context.Context, productId string) error {
	return models.NewDomainError("Not implemented yet")
}

type ProductsService interface {
	FetchProducts(ctx context.Context) (*ProductList, error)
	CreateProduct(ctx context.Context, porduct *models.Product) error
	DeleteProduct(ctx context.Context, productId string) error
}

// CreateProduct implements ProductsService.
func (p *productService) FetchProducts(ctx context.Context) (*ProductList, error) {
	fetchedProduct, err := p.productRepo.FetchProducts(ctx)
	if err != nil {
		return nil, err
	}

	var overviews ProductList
	for _, product := range fetchedProduct {
		overviews.Items = append(overviews.Items, product)
	}

	return &overviews, nil
}

func (p *productService) CreateProduct(ctx context.Context, product *models.Product) error {
	return p.productRepo.PersistProduct(ctx, product)
}

func NewProductService(productRepo repository.ProductRepository) ProductsService {
	return &productService{
		productRepo: productRepo,
	}
}
