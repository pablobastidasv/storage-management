package products

import (
	"github.com/gofiber/fiber/v2"
)

func HandleGetProducts(ctx *fiber.Ctx, storage ProductStorage) error {
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
