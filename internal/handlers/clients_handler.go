package handlers

import (
	"fmt"

	"co.bastriguez/inventory/internal/models"
	"co.bastriguez/inventory/internal/services"
	"github.com/gofiber/fiber/v2"
)

const (
	clientCreatedEvent string = "client-created"
)

type ClientHandlers struct {
	clientService services.ClientService
}

func NewClientHandler(clientService services.ClientService) *ClientHandlers {
	return &ClientHandlers{
		clientService: clientService,
	}
}

func (ch ClientHandlers) HandlePostClients(ctx *fiber.Ctx) error {
	req := struct {
		IdType   string `form:"id_type"`
		IdNumber string `form:"id_number"`
		Name     string `form:"name"`
	}{}
	ctx.BodyParser(&req)

	client := models.Client{
		Identity: models.Identity{
			DocumentType: models.DocumentType(req.IdType),
			Number:       req.IdNumber,
		},
		Name: req.Name,
	}

	if err := ch.clientService.Create(ctx.Context(), client); err != nil {
		return err
	}

	message := &AlertMessage{
		Message: "Client created successfully.",
        Level: Success,
	}
	ctx.Render("alert-messages", message)
	ctx.Response().Header.Add(
		hxRetarget,
		"#messages",
	)
	ctx.Response().Header.Add(
		hxTrigger,
		fmt.Sprintf("%s, %s", closeRightDrawerEvent, clientCreatedEvent),
	)
	return ctx.SendStatus(201)
}
