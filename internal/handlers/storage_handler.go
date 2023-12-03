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

type StorageHandlers struct {
	storageService  services.StorageService
	productsService services.ProductsService
}

func NewStorageHandler(service services.StorageService, productsService services.ProductsService) *StorageHandlers {
	return &StorageHandlers{
		storageService:  service,
		productsService: productsService,
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
	var err error

	indexVars["Products"], err = r.fetchStorageItems()
	if err != nil {
		return err
	}

	indexVars["Remissions"] = []RemissionItem{}

	return c.Render("index", indexVars)
}

func (r *StorageHandlers) AddProductFormHandler(c *fiber.Ctx) error {
	c.Response().Header.Add(hxTrigger, "open-right-drawer")
	products, err := r.loadProducts()
	if err != nil {
		return err
	}
	return c.Render("product_record_form", products)
}

func (r *StorageHandlers) loadProducts() ([]Product, error) {
	prods, err := r.productsService.FetchProducts()
	if err != nil {
		return nil, err
	}

	var products []Product
	for _, p := range prods.Items {
		products = append(products, Product{
			Id:           p.Id,
			Name:         p.Name,
			Presentation: translateUnit(p.Presentation),
		})
	}

	return products, nil
}

func (r *StorageHandlers) fetchStorageItems() ([]ProductItem, error) {
	storage, err := r.storageService.ItemsByStorage("unused")
	if err != nil {
		return nil, err
	}

	var products []ProductItem
	for _, i := range storage.Items {
		products = append(products, ProductItem{
			Id:     i.ProductId,
			Name:   i.Name,
			Amount: defineAmount(&i.Qty, i.Presentation),
		})
	}
	return products, nil

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
