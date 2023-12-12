package services

import (
	"co.bastriguez/inventory/internal/models"
	"co.bastriguez/inventory/internal/repository"
	"context"
	"fmt"
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
	AddProduct(ctx context.Context, storageId string, productId string, qty int) error
	ItemsByStorage(ctx context.Context, storageId string) (StorageItems, error)
}

type storageService struct {
	storageRepo repository.StorageRepository
	productRepo repository.ProductRepository
}

func (s storageService) ItemsByStorage(ctx context.Context, _ string) (StorageItems, error) {
	fetchedItems, err := s.storageRepo.FetchItemsByStorage(ctx, nil)
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

func (s storageService) AddProduct(ctx context.Context, _ string, productId string, qty int) error {
	if qty <= 0 {
		return NewWrongParameter("quantity MUST be more than zero (0)")
	}
	if len(productId) == 0 {
		return NewWrongParameter("product MUST be specified")
	}

	// there is only one storage therefore, its id should be retrieved
	storageId, err := s.fetchMainStorage(ctx)
	if err != nil {
		return err
	}

	item, err := s.storageRepo.FindItemByProductId(ctx, storageId, productId)
	if err != nil {
		return err
	}

	if item == nil {
		product, err := s.fetchProductById(ctx, productId)
		if err != nil {
			return err
		}

		item = &models.InventoryItem{
			Product: *product,
			Qty:     0,
		}
	}

	item.Qty += qty

	return s.storageRepo.UpsertItem(ctx, storageId, item)
}

func (s storageService) fetchMainStorage(ctx context.Context) (string, error) {
	storage, err := s.storageRepo.FindMainStorage(ctx)
	if err != nil {
		return "", err
	}

	return storage.Id, nil
}

func (s storageService) checkIfProductExist(ctx context.Context, productId string) error {
	productExist, err := s.productRepo.ExistProductById(ctx, productId)
	if err != nil {
		return err
	}
	if !productExist {
		return NewWrongParameter(
			fmt.Sprintf("the product with id %s was not found.", productId),
		)
	}
	return nil
}

func (s storageService) fetchProductById(ctx context.Context, productId string) (*models.InventoryProduct, error) {
	product, err := s.productRepo.FindProduct(ctx, productId)
	if err != nil {
		return nil, err
	}

	inventoryProduct := &models.InventoryProduct{
		Id:           product.Id,
		Name:         product.Name,
		Presentation: product.Presentation,
	}

	return inventoryProduct, err
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
