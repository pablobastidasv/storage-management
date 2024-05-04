package services

import (
	"context"

	"co.bastriguez/inventory/internal/models"
	"co.bastriguez/inventory/internal/repository"
)

type ClientService interface {
	Create(context.Context, models.Client) error
}

type clientService struct {
	repo repository.ClientRepository
}

func NewClientService(repository repository.ClientRepository) ClientService {
	return &clientService{
		repo: repository,
	}
}

func (c *clientService) Create(ctx context.Context, client models.Client) error {
	if err := validate(client); err != nil {
		return err
	}

	return c.repo.Save(ctx, client)
}

func validate(_ models.Client) error {
	return nil
}
