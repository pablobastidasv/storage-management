package services_test

import (
	"context"
	"fmt"
	"testing"

	"co.bastriguez/inventory/internal/models"
	"co.bastriguez/inventory/internal/services"
	"github.com/stretchr/testify/assert"
)

func Test_InvalidClient_ThenErrorIsReturned(t *testing.T) {
	tests := []struct {
		name            string
		client          models.Client
		expectedMessage string
	}{
		{
			name:            "given an empty document type",
			client:          aClient(WithDocType("")),
			expectedMessage: "Debes indicar el tipo de documento de identidad.",
		},
		{
			name:            "given an empty document number",
			client:          aClient(WithDocNumber("")),
			expectedMessage: "Debes indicar el número de documento de identidad.",
		},
		{
			name:            "given an empty client name",
			client:          aClient(WithName("")),
			expectedMessage: "Debes indicar el nombre.",
		},
		{
			name:            "given a blank client name",
			client:          aClient(WithName("   ")),
			expectedMessage: "Debes indicar el nombre.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sut := services.NewClientService(nil)
			err := sut.Create(context.Background(), tt.client)

			assert.Contains(t, tt.expectedMessage, err.Error())
		})
	}
}

type CaptureSave struct {
	client *models.Client
}

func (c *CaptureSave) SaveClient(_ context.Context, client models.Client) error {
	c.client = &client
	return nil
}

func Test_WhenNameIsTrim_NewNameHasNoSpaces(t *testing.T) {
	expectedName := "Jerry"
	client := aClient(WithName(fmt.Sprintf(" %s ", expectedName)))

	saver := &CaptureSave{}

	sut := services.NewClientService(saver)
	sut.Create(context.Background(), client)

	assert.Equal(t, expectedName, saver.client.Name)
}

type ProcClient func(client *models.Client)

func WithDocNumber(docNumber string) ProcClient {
	return func(client *models.Client) {
		client.Identity.Number = docNumber
	}
}

func WithName(name string) ProcClient {
	return func(client *models.Client) {
		client.Name = name
	}
}

func WithDocType(docType models.DocumentType) ProcClient {
	return func(client *models.Client) {
		client.Identity.DocumentType = docType
	}
}

func aClient(proc ...ProcClient) models.Client {
	c := models.Client{
		Identity: models.Identity{
			DocumentType: models.NIT,
			Number:       "1123123123",
		},
		Name: "Pachito Eche",
	}

	for _, p := range proc {
		p(&c)
	}

	return c
}
