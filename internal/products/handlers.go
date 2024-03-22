package products

import (
	"github.com/gofiber/fiber/v2"
)

var storage ProductStorage

func InitHandlers(s ProductStorage) {
    storage = s
}

func HandleGetProducts(ctx *fiber.Ctx) error {
	products, err := storage.ListProducts(ctx.Context())
	if err != nil {
		return err
	}

	productDetails := []struct {
		Name string
	}{}

	for _, p := range products {
		prod := struct {
			Name string
		}{
			Name: p.name,
		}
		productDetails = append(productDetails, prod)
	}

	productsView := struct {
		Products []struct {
			Name string
		}
	}{
        Products: productDetails,
    }

	return ctx.Render("PoductList", productsView)

}
