package creating

import (
	"fmt"

	"co.bastriguez/inventory/kit/command"
	"golang.org/x/net/context"
)

const CreateCommandType command.Type = "command.creating.client"

type CreateClientCommand struct {
	docType     string
	docNumber   string
	name        string
	address     string
	email       string
	phoneNumber string
}

func NewCreateClientCommand(docType, docNumber, name, address, email, phoneNumber string) CreateClientCommand {
	return CreateClientCommand{
		docType:     docType,
		docNumber:   docNumber,
		name:        name,
		address:     address,
		email:       email,
		phoneNumber: phoneNumber,
	}
}

func (c CreateClientCommand) Type() command.Type {
	return CreateCommandType
}

type CreateClientCommandHandler struct {
	service ClientCreator
}

func NewCreateClientCommandHandler(service ClientCreator) *CreateClientCommandHandler {
	return &CreateClientCommandHandler{
		service: service,
	}
}

func (h CreateClientCommandHandler) Handle(ctx context.Context, cmd command.Command) error {
	createClientCommand, ok := cmd.(CreateClientCommand)
	if !ok {
		return fmt.Errorf("unexpected command")
	}

	return h.service.Create(
		ctx,
		createClientCommand.docType,
		createClientCommand.docNumber,
		createClientCommand.name,
        createClientCommand.address,
        createClientCommand.email,
        createClientCommand.phoneNumber,
	)
}
