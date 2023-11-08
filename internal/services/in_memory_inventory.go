package services

import "co.bastriguez/inventory/internal/entities"

type InMemoryInventoryService struct {
	storage    entities.Storage
	remissions []entities.Remission
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

func (m InMemoryInventoryService) RetrieveOpenRemissions() []entities.Remission {
	m.remissions = []entities.Remission{
		{
			Id: "A",
			Client: entities.Client{
				Name: "Client A",
			},
			Product: entities.Product{
				Name:         "Product W",
				Presentation: entities.KG,
			},
			Qty: 15,
		},
		{
			Id: "B",
			Client: entities.Client{
				Name: "Client B",
			},
			Product: entities.Product{
				Name:         "Product Z",
				Presentation: entities.Grms,
			},
			Qty: 24,
		},
	}

	return m.remissions
}
