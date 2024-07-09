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
	IdType   string `form:"id_type"`
	IdNumber string `form:"id_number"`
	Name     string `form:"name"`
}

func PostClientHandler(bus command.Bus) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req createRequest
		c.BodyParser(&req)

		cmd := creating.NewCreateClientCommand(req.IdType, req.IdNumber, req.Name)
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
