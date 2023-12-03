package services

import (
	"co.bastriguez/inventory/internal/models"
	"co.bastriguez/inventory/internal/repository"
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

func (p productService) FetchProducts() (ProductList, error) {
	fetchedProduct, err := p.productRepo.FetchProducts()
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
	FetchProducts() (ProductList, error)
}

func NewProductService(productRepo repository.ProductRepository) ProductsService {
	return &productService{
		productRepo: productRepo,
	}
}
