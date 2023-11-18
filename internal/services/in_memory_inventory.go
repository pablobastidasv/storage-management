package services

import (
	"co.bastriguez/inventory/internal/models"
	"time"
)

type InMemoryInventoryService struct {
	storage    models.Storage
	remissions []models.Remission
}

func NewInMemoryInventoryService() InventoryService {
	service := &InMemoryInventoryService{}

	service.storage = models.Storage{
		Content: []models.InventoryContent{
			{
				Product: models.Product{
					Name: "Product A",
				},
				Qty: 15,
			},
			{
				Product: models.Product{
					Name: "Product B",
				},
				Qty: 15,
			},
			{
				Product: models.Product{
					Name: "Product C",
				},
				Qty: 15,
			},
		},
	}
	service.remissions = []models.Remission{
		{
			Id: "A",
			Client: models.Client{
				Name: "Client A",
			},
			Product: models.Product{
				Name:         "Product W",
				Presentation: models.KG,
			},
			Qty:       15,
			CreatedAt: time.Now(),
		},
		{
			Id: "B",
			Client: models.Client{
				Name: "Client B",
			},
			Product: models.Product{
				Name:         "Product Z",
				Presentation: models.Grms,
			},
			Qty:       24,
			CreatedAt: time.Now(),
		},
	}

	return service
}

func (m InMemoryInventoryService) RetrieveProducts() []models.InventoryContent {
	return m.storage.Content
}

func (m InMemoryInventoryService) RetrieveOpenRemissions() []models.Remission {
	return m.remissions
}
