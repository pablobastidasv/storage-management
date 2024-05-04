package handlers

import (
	"github.com/gofiber/fiber/v2"

	"co.bastriguez/inventory/internal/models"
)

type Option struct {
	Id    string
	Label string
}

type AdminHandlers struct{}

func (a *AdminHandlers) HandleAdminClientsPage(c *fiber.Ctx) error {
	return c.Render("pages/admin-clients", nil, "admin-template")
}

func (a *AdminHandlers) HandleAdminCreateClientFormFragment(c *fiber.Ctx) error {
	params := struct {
	}{}

	c.Response().Header.Add(hxTrigger, openRightDrawerEvent)
	return c.Render("admin_clients_add_form", params)
}

func (a *AdminHandlers) HandleAdminProductsPage(c *fiber.Ctx) error {
	return c.Render("pages/admin-products", nil, "admin-template")
}

func (a *AdminHandlers) HandleAdminHomePage(c *fiber.Ctx) error {
	return c.Render("pages/admin", nil, "admin-template")
}

func (a *AdminHandlers) HandleAdminCreateProductFormFragment(c *fiber.Ctx) error {
	params := struct {
		Presentations []Option
	}{}

	for _, p := range models.ListPresentations() {
		params.Presentations = append(params.Presentations, Option{
			Id:    string(p),
			Label: translateUnit(p),
		})
	}

	c.Response().Header.Add(hxTrigger, openRightDrawerEvent)
	return c.Render("admin_products_add_form", params)
}

func NewAdminHandler() *AdminHandlers {
	return &AdminHandlers{}
}
