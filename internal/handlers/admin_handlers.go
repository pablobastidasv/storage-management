package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type AdminHandlers struct {
}

func (a *AdminHandlers) HandleAdminHomePage(c *fiber.Ctx) error {
	params := struct {
		Products []Product
	}{
		Products: []Product{
			{
				Id:           "The ID",
				Name:         "This is a product name",
				Presentation: "Gramos",
			}, {
				Id:           "Another",
				Name:         "Burrito",
				Presentation: "Unidades",
			},
		},
	}
	return c.Render("pages/admin", params, "general-template")
}

func (a *AdminHandlers) HandleAdminCreateProductFormFragment(c *fiber.Ctx) error {
	c.Response().Header.Add(hxTrigger, openRightDrawerEvent)
	return c.Render("admin_products_add_form", nil)
}

func NewAdminHandler() *AdminHandlers {
	return &AdminHandlers{}
}
