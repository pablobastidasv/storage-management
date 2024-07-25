package creating_test

import (
	"context"
	"errors"
	"testing"

	"co.bastriguez/inventory/internal/creating"
	"co.bastriguez/inventory/internal/inventory"
	"co.bastriguez/inventory/internal/platform/storage/storagemocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_ClientService_CreateClient_Error(t *testing.T) {
	docType := "CC"
	docNumber := "1234567890"
	name := "Pepe Perez"

	repomock := new(storagemocks.ClientPersister)
	repomock.On("Save", mock.Anything, mock.Anything).Return(errors.New("there was an unexpected error"))

	sut := creating.NewClientService(repomock)
	err := sut.Create(context.Background(), docType, docNumber, name, "", "", "")

	assert.Error(t, err)
}

func Test_ClientService_CreateClient_Succeed(t *testing.T) {
	docType := "CC"
	docNumber := "1234567890"
	name := "Pepe Perez"
    address := "Address 15"
    email := "email@email.test"
    phone := "0123456789"

    expectedClient := inventory.Client{
    	Document:          inventory.Document{
    		DocumentType: inventory.DocumentType(docType),
    		Number:       inventory.DocumentNumber(docNumber),
    	},
    	Name:              inventory.ClientName(name),
    	Address:           inventory.ClientAddress(address),
    	Email:             inventory.ClientEmail(email),
    	PhoneNumber: inventory.ClientPhoneNumber(phone),
    }

	repomock := new(storagemocks.ClientPersister)
	repomock.On("Save", mock.Anything, expectedClient).Return(nil)

    sut := creating.NewClientService(repomock)
    err := sut.Create(context.Background(), docType, docNumber, name, address, email, phone)

    assert.NoError(t, err)
    repomock.AssertExpectations(t)
}
