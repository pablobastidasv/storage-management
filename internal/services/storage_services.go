package services

import (
	"co.bastriguez/inventory/internal/models"
	"co.bastriguez/inventory/internal/repository"
	"fmt"
	"os/exec"
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

func (s storageService) AddProduct(_ string, productId string, qty int) error {
	// there is only one storage therefore, its id should be retrieved
	storageId, err := s.fetchMainStorage()
	if err != nil {
		return err
	}

	if err := s.checkIfProductExist(productId); err != nil {
		return err
	}

	item, err := s.storageRepo.FindItemBy(storageId, productId)
	if err != nil {
		return err
	}

	if item == nil {
		item = &models.InventoryItem{
			Product: models.Product{
				Id: productId,
			},
			Qty: 0,
		}
	}

	item.Qty += qty

	return s.storageRepo.UpdateItem(storageId, item)
}

func (s storageService) fetchMainStorage() (string, error) {
	storage, err := s.storageRepo.FindMainStorage()
	if err != nil {
		return "", err
	}

	return storage.Id, nil
}

func (s storageService) checkIfProductExist(productId string) error {
	productExist, err := s.productRepo.ExistProductById(productId)
	if err != nil {
		return err
	}
	if !productExist {
		return &exec.Error{
			Name: fmt.Sprintf("the product with id %s was not found.", productId),
		}
	}
	return nil
}

func NewStorageService(
	storageRepo repository.StorageRepository,
	productRepo repository.ProductRepository,
) StorageService {
	return &storageService{
		storageRepo: storageRepo,
		productRepo: productRepo,
	}
}
