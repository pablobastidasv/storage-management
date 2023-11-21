package handlers

import (
	"co.bastriguez/inventory/internal/models"
	"co.bastriguez/inventory/internal/services"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"time"
)

const (
	hxTrigger = "HX-Trigger"
)

type (
	StorageHandler struct {
		inventoryService services.InventoryService
	}

	Product struct {
		Id           string
		Name         string
		Presentation string
	}

	ProductItem struct {
		Id     string
		Name   string
		Amount string
	}

	RemissionItem struct {
		ClientName string
		ProductItem
		CreatedAt time.Time
	}
)

func NewStorageHandler(service services.InventoryService) *StorageHandler {
	return &StorageHandler{
		inventoryService: service,
	}
}

func (r *StorageHandler) GetProductsHandler(ctx *fiber.Ctx) error {
	productItems := r.generateProductItems()
	return ctx.Render("inventory", productItems)
}

func (r *StorageHandler) generateProductItems() []ProductItem {
	var productItems []ProductItem

	var products = r.inventoryService.RetrieveProducts()
	for _, product := range products {
		item := ProductItem{
			Id:     product.Product.Id,
			Name:   product.Product.Name,
			Amount: defineAmount(&product.Qty, product.Product.Presentation),
		}
		productItems = append(productItems, item)
	}
	return productItems
}

func (r *StorageHandler) PutProductsHandler(ctx *fiber.Ctx) error {
	ctx.Response().Header.Add(hxTrigger, "close-right-drawer, load-storage-products")
	return ctx.SendStatus(204)
}

func (r *StorageHandler) StorageRemissionsHandler(ctx *fiber.Ctx) error {
	remissions := r.generateRemissionItems()
	return ctx.Render("remissions", remissions)
}

func (r *StorageHandler) generateRemissionItems() []RemissionItem {
	remissions := r.inventoryService.RetrieveOpenRemissions()
	var remissionItems []RemissionItem
	for _, remission := range remissions {
		item := RemissionItem{
			ClientName: remission.Client.Name,
			ProductItem: ProductItem{
				Id:     remission.Product.Id,
				Name:   remission.Product.Name,
				Amount: defineAmount(&remission.Qty, remission.Product.Presentation),
			},
			CreatedAt: remission.CreatedAt,
		}
		remissionItems = append(remissionItems, item)
	}
	return remissionItems
}

func (r *StorageHandler) InventoryHomePageHandler(c *fiber.Ctx) error {
	indexVars := make(map[string]interface{})
	indexVars["Products"] = r.generateProductItems()
	indexVars["Remissions"] = r.generateRemissionItems()

	return c.Render("index", indexVars)
}

func (r *StorageHandler) AddProductFormHandler(c *fiber.Ctx) error {
	c.Response().Header.Add(hxTrigger, "open-right-drawer")

	products := []Product{
		{
			Id:           "p-a",
			Name:         "Product A",
			Presentation: translateUnit(models.KG),
		}, {
			Id:           "p-b",
			Name:         "Product B",
			Presentation: translateUnit(models.Grms),
		}, {
			Id:           "p-c",
			Name:         "Product C",
			Presentation: translateUnit(models.Amount),
		},
	}

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
