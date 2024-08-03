package inventory

import (
	"context"
	"fmt"
	"net/mail"
	"strings"

	"co.bastriguez/inventory/internal/inventory/validators"
)

// root error used to typify an error in order to control it in the client response
type ValidationError struct {
	Message string
	Details []FieldError
}

func (v *ValidationError) Error() string {
	return v.Message
}

func (v *ValidationError) AddDetail(field, message string) {
	v.Details = append(v.Details, FieldError{
		Field:    field,
		Messsage: message,
	})
}

type FieldError struct {
	Field    string
	Messsage string
}

func (f *FieldError) Error() string {
	return fmt.Sprintf("%s: %s", f.Field, f.Messsage)
}

func NewFieldError(field, message string) *FieldError {
	return &FieldError{
		Field:    field,
		Messsage: message,
	}
}

type (
	DocumentType int

	DocumentNumber    string
	ClientName        string
	ClientAddress     string
	ClientEmail       string
	ClientPhoneNumber string

	Client struct {
		Document    Document
		Name        ClientName
		Address     ClientAddress
		Email       ClientEmail
		PhoneNumber ClientPhoneNumber
	}

	Document struct {
		DocumentType DocumentType
		Number       DocumentNumber
	}
)

const (
	DocTypeCC DocumentType = iota
	DocTypeNit
	DocTypeCE

	fieldDocType     string = "doc_type"
	fieldDocNumber   string = "doc_number"
	fieldEmail       string = "email"
	fieldName        string = "name"
	fieldPhoneNumber string = "phone_number"
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
	if strings.Trim(documentType, " ") == "" {
		return DocumentType(-1), NewFieldError(fieldDocType, "no puede estar vacio")
	}

	switch documentType {
	case "CC":
		return DocTypeCC, nil
	case "NIT":
		return DocTypeNit, nil
	case "CE":
		return DocTypeCE, nil
	default:
		msg := fmt.Sprintf("document type '%s' is not valid", documentType)
		return DocumentType(-1), NewFieldError(fieldDocType, msg)
	}
}

func NewClientName(name string) (ClientName, error) {
	if strings.Trim(name, " ") == "" {
		return ClientName(""), NewFieldError(fieldName, "no puede estar vacio")
	}
	return ClientName(name), nil
}

func NewDocumentNumber(documentNumber string) (DocumentNumber, error) {
	if strings.Trim(documentNumber, " ") == "" {
		return DocumentNumber(""), NewFieldError(fieldDocNumber, "no puede estar vacio")
	}

	if !validators.IsDocumentNumber(documentNumber) {
		expectedErrorMessage := fmt.Sprintf("numero de documento '%s' no es valido", documentNumber)
		return DocumentNumber(""), NewFieldError(fieldDocNumber, expectedErrorMessage)
	}

	return DocumentNumber(documentNumber), nil
}

func NewClientAddress(address string) (ClientAddress, error) {
	return ClientAddress(address), nil
}

func NewClientPhone(phone string) (ClientPhoneNumber, error) {
	if phone != "" && !validators.IsPhoneNumber(phone) {
		expectedErrorMessage := fmt.Sprintf("el numero de telefono/celular dado no es valido")
		return ClientPhoneNumber(""), NewFieldError(fieldPhoneNumber, expectedErrorMessage)
	}

	return ClientPhoneNumber(phone), nil
}

func NewClientEmail(email string) (ClientEmail, error) {
	if strings.Trim(email, " ") == "" {
		return ClientEmail(""), NewFieldError(fieldEmail, "no puede estar vacio")
	}

	if _, err := mail.ParseAddress(email); err != nil {
		return ClientEmail(""), NewFieldError(fieldEmail, "formato de email no es valido")
	}

	return ClientEmail(email), nil
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

	clientAddress, err := NewClientAddress(address)
	if err != nil {
		return Client{}, err
	}

	clientPhone, err := NewClientPhone(phoneNumber)
	if err != nil {
		return Client{}, err
	}

	clientEmail, err := NewClientEmail(email)
	if err != nil {
		return Client{}, err
	}

	return Client{
		Document: Document{
			DocumentType: documentType,
			Number:       documentNumber,
		},
		Name:        clientName,
		Address:     clientAddress,
		Email:       clientEmail,
		PhoneNumber: clientPhone,
	}, nil
}
