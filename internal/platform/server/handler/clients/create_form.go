package clients

import (
	"co.bastriguez/inventory/internal/platform/server/views"
	"co.bastriguez/inventory/internal/platform/server/views/pages"
	"github.com/gofiber/fiber/v2"
)

func GetClientFormHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return views.Render(c, pages.CreateClient())
	}
}
