package routes

import "github.com/gofiber/fiber/v2"

type Routes interface {
	DefineRoutes(mainRoute fiber.Router)
}

func New() Routes {
	return &inventoryRoutes{}
}

type inventoryRoutes struct{}

func (r inventoryRoutes) DefineRoutes(storages fiber.Router) {
	storages.Get("/main/products", StorageProductsHandler)
	storages.Get("/main/remissions", StorageRemissionsHandler)
}

func StorageProductsHandler(ctx *fiber.Ctx) error {
	return ctx.Render("inventory", nil)
}

func StorageRemissionsHandler(ctx *fiber.Ctx) error {
	return ctx.Render("remissions", nil)
}
