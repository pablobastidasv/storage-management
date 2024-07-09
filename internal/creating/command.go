package creating

import (
	"fmt"

	"co.bastriguez/inventory/kit/command"
	"golang.org/x/net/context"
)

const CreateCommandType command.Type = "command.creating.client"

type CreateClientCommand struct {
	docType   string
	docNumber string
	name      string
}

func NewCreateClientCommand(docType, docNumber, name string) CreateClientCommand {
	return CreateClientCommand{
		docType:   docType,
		docNumber: docNumber,
		name:      name,
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
	)
}
