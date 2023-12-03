package handlers

import (
	"co.bastriguez/inventory/internal/services"
	"github.com/gofiber/fiber/v2"
)

type ProductHandlers struct {
	service services.ProductsService
}

func (p *ProductHandlers) HandleGetProducts(ctx *fiber.Ctx) error {
	return ctx.SendString("Hola from Products.")
}

func NewProductHandler(service services.ProductsService) *ProductHandlers {
	return &ProductHandlers{
		service: service,
	}
}
