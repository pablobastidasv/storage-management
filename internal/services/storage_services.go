package services

import (
	"co.bastriguez/inventory/internal/models"
	"co.bastriguez/inventory/internal/repository"
)

type (
	DetailStorageItem struct {
		ProductId    string
		Name         string
		Presentation models.Presentation
		Qty          int
	}

	StorageItems struct {
		Items []DetailStorageItem
	}
)

type StorageService interface {
	AddProduct(storageId string, productId string, qty int) error
	RemoveProduct(storageId string, productId string, qty int) error
	ItemsByStorage(storageId string) (StorageItems, error)
}

type storageService struct {
	storageRepo repository.StorageRepository
	productRepo repository.ProductRepository
}

func (s storageService) ItemsByStorage(_ string) (StorageItems, error) {
	fetchedItems, err := s.storageRepo.FetchItemsByStorage(nil)
	if err != nil {
		return StorageItems{}, err
	}

	var items []DetailStorageItem
	for _, i := range fetchedItems {
		items = append(items, DetailStorageItem{
			ProductId:    i.Product.Id,
			Name:         i.Product.Name,
			Presentation: i.Product.Presentation,
			Qty:          i.Qty,
		})
	}

	return StorageItems{
		Items: items,
	}, nil
}

func (s storageService) RemoveProduct(storageId string, productId string, qty int) error {
	//TODO implement me
	panic("implement me")
}

func (s storageService) AddProduct(storageId string, productId string, qty int) error {
	//TODO implement me
	panic("implement me")
}

func NewStorage(storageRepo repository.StorageRepository, productRepo repository.ProductRepository) StorageService {
	return &storageService{
		storageRepo: storageRepo,
		productRepo: productRepo,
	}
}
