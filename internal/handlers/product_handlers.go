package handlers

import (
	"co.bastriguez/inventory/internal/services"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type ProductHandlers struct {
	service services.ProductsService
}

func (p *ProductHandlers) HandlePostProducts(c *fiber.Ctx) error {
	c.Response().Header.Add(hxTrigger, fmt.Sprintf("%s, %s", closeRightDrawerEvent, productCreatedEvent))
	return c.SendStatus(201)
}

func (p *ProductHandlers) HandleGetProducts(c *fiber.Ctx) error {
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
			}, {
				Id:           "42",
				Name:         "Wata",
				Presentation: "Ltrs",
			}, {
				Id:           "24",
				Name:         "ataw",
				Presentation: "Grms",
			},
		},
	}
	return c.Render("products-list", params)
}

func NewProductHandler(service services.ProductsService) *ProductHandlers {
	return &ProductHandlers{
		service: service,
	}
}
