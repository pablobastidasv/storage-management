package clients

import (
	"errors"
	"net/http"

	"co.bastriguez/inventory/internal/creating"
	"co.bastriguez/inventory/internal/inventory"
	"co.bastriguez/inventory/kit/command"
	"github.com/gofiber/fiber/v2"
)

type createRequest struct {
	DocType     string `form:"doc_type"`
	IdNumber    string `form:"doc_number"`
	Name        string `form:"name"`
	Address     string `form:"address"`
	Email       string `form:"email"`
	PhoneNumber string `form:"phone_number"`
}

func PostClientHandler(bus command.Bus) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req createRequest
		c.BodyParser(&req)

		cmd := creating.NewCreateClientCommand(req.DocType, req.IdNumber, req.Name, req.Address, req.Email, req.PhoneNumber)
		err := bus.DispatchCommand(c.Context(), cmd)
		if err != nil {
			// TODO: Check how to have a global exception handler for
			//       error on entities when they are process (422)
			switch {
			case errors.Is(err, inventory.ErrInvalidDocumentType),
				errors.Is(err, inventory.ErrInvalidDocumentNumber),
				errors.Is(err, inventory.ErrInvalidClientName):
				c.Response().SetStatusCode(http.StatusUnprocessableEntity)
				return nil
			default:
				return err
			}
		}

		c.Response().SetStatusCode(http.StatusCreated)
		return nil
	}
}
