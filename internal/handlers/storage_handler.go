package handlers

import (
	"co.bastriguez/inventory/internal/models"
	"co.bastriguez/inventory/internal/services"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

const (
	hxTrigger = "HX-Trigger"
)

type (
	StorageHandlers struct {
		storageService services.StorageService
	}
)

func NewStorage(service services.StorageService) *StorageHandlers {
	return &StorageHandlers{
		storageService: service,
	}
}

func (r *StorageHandlers) GetProductsHandler(ctx *fiber.Ctx) error {
	var productItems []ProductItem
	return ctx.Render("inventory", productItems)
}

func (r *StorageHandlers) PutProductsHandler(ctx *fiber.Ctx) error {
	ctx.Response().Header.Add(hxTrigger, "close-right-drawer, load-repository-products")
	return ctx.SendStatus(204)
}

func (r *StorageHandlers) StorageRemissionsHandler(ctx *fiber.Ctx) error {
	var remissions []RemissionItem
	return ctx.Render("remissions", remissions)
}

func (r *StorageHandlers) InventoryHomePageHandler(c *fiber.Ctx) error {
	indexVars := make(map[string]interface{})
	indexVars["Products"] = []ProductItem{}
	indexVars["Remissions"] = []RemissionItem{}

	return c.Render("index", indexVars)
}

func (r *StorageHandlers) AddProductFormHandler(c *fiber.Ctx) error {
	c.Response().Header.Add(hxTrigger, "open-right-drawer")
	var products []Product
	return c.Render("product_record_form", products)
}

func defineAmount(qty *int, presentation models.Presentation) string {
	return fmt.Sprintf("%d %s", *qty, translateUnit(presentation))
}

func translateUnit(unit models.Presentation) string {
	switch unit {
	case models.KG:
		return "Kilogramos"
	case models.Amount:
		return "Cantidad"
	case models.Grms:
		return "Gramos"
	default:
		return "Desconocido"
	}
}
