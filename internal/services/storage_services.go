package services

import "co.bastriguez/inventory/internal/repository"

type StorageService interface {
	AddProduct(storageId string, productId string, qty int) error
	RemoveProduct(storageId string, productId string, qty int) error
}

type storageService struct {
	storageRepo repository.StorageRepository
	productRepo repository.ProductRepository
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
