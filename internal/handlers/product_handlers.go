package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"co.bastriguez/inventory/internal/models"
	"co.bastriguez/inventory/internal/services"
)

type ProductHandlers struct {
	service services.ProductsService
}

func (p *ProductHandlers) HandlePostProducts(c *fiber.Ctx) error {
	req := struct {
		Id           string
		Name         string
		Presentation string
	}{
		Id: uuid.NewString(),
	}
	c.BodyParser(&req)

	product, err := models.CreateProduct(req.Id, req.Name, req.Presentation)
	if err != nil {
		return err
	}
	p.service.CreateProduct(c.Context(), product)

	c.Response().Header.Add(
		hxTrigger,
		fmt.Sprintf("%s, %s", closeRightDrawerEvent, productCreatedEvent),
	)
	return c.SendStatus(201)
}

func (p *ProductHandlers) HandleGetProducts(c *fiber.Ctx) error {
	productList, err := p.service.FetchProducts(c.Context())
	if err != nil {
		return err
	}

	params := struct {
		Products []Product
	}{}

	for _, p := range productList.Items {
		params.Products = append(params.Products, Product{
			Id:           p.Id.ToString(),
			Name:         p.Name.ToString(),
			Presentation: translateUnit(p.Presentation),
		})
	}

	return c.Render("products-list", params)
}

func NewProductHandler(service services.ProductsService) *ProductHandlers {
	return &ProductHandlers{
		service: service,
	}
}
