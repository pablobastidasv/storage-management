package creating

import (
	"context"

	"co.bastriguez/inventory/internal/inventory"
)

type clientService struct {
	repo inventory.ClientRepository
}

func NewClientService(repo inventory.ClientRepository) ClientCreator {
	return &clientService{
		repo: repo,
	}
}

func (c *clientService) Create(ctx context.Context, docType, docNumber, name string) error {
    client, err := inventory.NewClient(docType, docNumber, name)
    if err != nil {
        return err
    }

    return c.repo.Save(ctx, client)
}
