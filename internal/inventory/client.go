package inventory

import (
	"context"
	"errors"
	"fmt"
)

var ErrInvalidDocumentType = errors.New("invalid document type")
var ErrInvalidDocumentNumber = errors.New("invalid document number")
var ErrInvalidClientName = errors.New("invalid client name")

type (
	DocumentType      string
	DocumentNumber    string
	ClientName        string
	ClientAddress     string
	ClientEmail       string
	ClientPhoneNumber string

	Client struct {
		Document          Document
		Name              ClientName
		Address           ClientAddress
		Email             ClientEmail
		ClientPhoneNumber ClientPhoneNumber
	}

	Document struct {
		DocumentType DocumentType
		Number       DocumentNumber
	}
)

//go:generate mockery --case=snake --outpkg=storagemocks --output=../platform/storage/storagemocks --name ClientPersister
type (
	ClientRepository interface {
		ClientPersister
	}

	ClientPersister interface {
		Save(context.Context, Client) error
	}
)

func NewDocumentType(documentType string) (DocumentType, error) {
	if documentType == "" {
		return DocumentType(""), fmt.Errorf("%w, %s", ErrInvalidDocumentType, documentType)
	}

	return DocumentType(documentType), nil
}

func NewClientName(name string) (ClientName, error) {
	if name == "" {
		return ClientName(""), fmt.Errorf("%w, %s", ErrInvalidClientName, name)
	}
	return ClientName(name), nil
}

func NewDocumentNumber(documentNumber string) (DocumentNumber, error) {
	return DocumentNumber(documentNumber), nil
}

func NewClient(docType, docNumber, name, address, email, phoneNumber string) (Client, error) {
	documentType, err := NewDocumentType(docType)
	if err != nil {
		return Client{}, err
	}

	documentNumber, err := NewDocumentNumber(docNumber)
	if err != nil {
		return Client{}, err
	}

	clientName, err := NewClientName(name)
	if err != nil {
		return Client{}, err
	}

	return Client{
		Document: Document{
			DocumentType: documentType,
			Number:       documentNumber,
		},
		Name: clientName,
	}, nil
}
