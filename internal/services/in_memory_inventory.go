package services

import "co.bastriguez/inventory/internal/entities"

type InMemoryInventoryService struct {
	storage entities.Storage
}

func NewInMemoryInventoryService() InventoryService {
	service := &InMemoryInventoryService{}

	service.storage = entities.Storage{
		Content: []entities.InventoryContent{
			{
				Product: entities.Product{
					Name: "Product A",
				},
				Qty: 15,
			},
			{
				Product: entities.Product{
					Name: "Product B",
				},
				Qty: 15,
			},
			{
				Product: entities.Product{
					Name: "Product C",
				},
				Qty: 15,
			},
		},
	}
	return service
}

func (m InMemoryInventoryService) RetrieveInventory() []entities.InventoryContent {
	return m.storage.Content
}
