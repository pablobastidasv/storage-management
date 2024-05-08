package services

import (
	"context"
	"strings"

	"co.bastriguez/inventory/internal/models"
	"co.bastriguez/inventory/internal/repository"
)

type (
	ClientService interface {
		Create(context.Context, models.Client) error
	}
	clientSaver interface {
		SaveClient(context.Context, models.Client) error
	}

	clientService struct {
		repo repository.ClientRepository
	}
)

func NewClientService(repository clientSaver) ClientService {
	return &clientService{
		repo: repository,
	}
}

func (c *clientService) Create(ctx context.Context, client models.Client) error {
	if err := validate(&client); err != nil {
		return err
	}

	return c.repo.SaveClient(ctx, client)
}

func validate(client *models.Client) error {
	var messages []string
	if client.Identity.DocumentType == "" {
		messages = append(messages, "Debes indicar el tipo de documento de identidad.")
	}
	if client.Identity.Number == "" {
        messages = append(messages, "Debes indicar el número de documento de identidad.")
	}
	client.Name = strings.Trim(client.Name, " ")
	if client.Name == "" {
        messages = append(messages, "Debes indicar el nombre.")
	}

    if len(messages) > 0{
        return models.NewValidationError(messages)
    }

	return nil
}
