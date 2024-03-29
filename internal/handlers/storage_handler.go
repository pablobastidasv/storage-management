package handlers

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"

	"co.bastriguez/inventory/internal/models"
	"co.bastriguez/inventory/internal/services"
)

type StorageHandlers struct {
	storageService  services.StorageService
	productsService services.ProductsService
}

func NewStorageHandler(
	service services.StorageService,
	productsService services.ProductsService,
) *StorageHandlers {
	return &StorageHandlers{
		storageService:  service,
		productsService: productsService,
	}
}

func (r *StorageHandlers) HandleGetItems(ctx *fiber.Ctx) error {
	productItems, err := r.fetchStorageItems(ctx.Context())
	if err != nil {
		return err
	}

	return ctx.Render("inventory", productItems)
}

func (r *StorageHandlers) HandlePutProducts(ctx *fiber.Ctx) error {
	var request PutProductsRequest
	var err error
	err = ctx.BodyParser(&request)
	if err != nil {
		return err
	}

	err = r.storageService.AddProduct(ctx.Context(), "main", request.Product, request.Qty)
	if err != nil {
		return err
	}

	ctx.Response().Header.Add(
		hxTrigger,
		fmt.Sprintf("%s, load-storage-products", closeRightDrawerEvent),
	)
	return ctx.SendStatus(204)
}

func (r *StorageHandlers) HandleGetRemissions(ctx *fiber.Ctx) error {
	var remissions []RemissionItem
	return ctx.Render("remissions", remissions)
}

func (r *StorageHandlers) HandleInventoryHomePage(c *fiber.Ctx) error {
	indexVars := make(map[string]interface{})
	var err error

	indexVars["Products"], err = r.fetchStorageItems(c.Context())
	if err != nil {
		return err
	}

	indexVars["Remissions"] = []RemissionItem{}

	return c.Render("pages/storage-management", indexVars, "general-template")
}

func (r *StorageHandlers) HandleAddProductFormFragment(c *fiber.Ctx) error {
	c.Response().Header.Add(hxTrigger, openRightDrawerEvent)
	products, err := r.loadProducts(c.Context())
	if err != nil {
		return err
	}
	return c.Render("product_record_form", products)
}

// TODO: Remove this one from here
func (r *StorageHandlers) loadProducts(ctx context.Context) ([]Product, error) {
	prods, err := r.productsService.FetchProducts(ctx)
	if err != nil {
		return nil, err
	}

	var products []Product
	for _, p := range prods.Items {
		products = append(products, Product{
			Id:           p.Id.ToString(),
			Name:         p.Name.ToString(),
			Presentation: translateUnit(p.Presentation),
		})
	}

	return products, nil
}

func (r *StorageHandlers) fetchStorageItems(ctx context.Context) ([]ProductItem, error) {
	storage, err := r.storageService.ItemsByStorage(ctx, "unused")
	if err != nil {
		return nil, err
	}

	var products []ProductItem
	for _, i := range storage.Items {
		products = append(products, ProductItem{
			Id:           i.ProductId,
			Name:         i.Name,
			Amount:       defineAmount(&i.Qty, i.Presentation),
			Qty:          i.Qty,
			Presentation: translateUnit(i.Presentation),
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
		return "Unidades"
	case models.Grms:
		return "Gramos"
	default:
		return "Desconocido"
	}
}
