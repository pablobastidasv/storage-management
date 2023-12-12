package services

import (
	"co.bastriguez/inventory/internal/models"
	"co.bastriguez/inventory/internal/repository"
	"context"
)

type (
	productService struct {
		productRepo repository.ProductRepository
	}

	ProductList struct {
		Items []ProductOverview
	}

	ProductOverview struct {
		Id           string
		Name         string
		Presentation models.Presentation
	}
)

func (p productService) FetchProducts(ctx context.Context) (ProductList, error) {
	fetchedProduct, err := p.productRepo.FetchProducts(ctx)
	if err != nil {
		return ProductList{}, err
	}

	var overviews ProductList
	for _, product := range fetchedProduct {
		overview := ProductOverview{
			Id:           product.Id,
			Name:         product.Name,
			Presentation: product.Presentation,
		}
		overviews.Items = append(overviews.Items, overview)
	}

	return overviews, nil
}

type ProductsService interface {
	FetchProducts(ctx context.Context) (ProductList, error)
}

func NewProductService(productRepo repository.ProductRepository) ProductsService {
	return &productService{
		productRepo: productRepo,
	}
}
