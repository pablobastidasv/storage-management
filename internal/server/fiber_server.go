package server

import (
	"co.bastriguez/inventory/internal/routes"
	"co.bastriguez/inventory/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

type (
	fiberServer struct{}
)

func (f fiberServer) Start(addr string) error {
	// Initialize standard Go html template engine
	engine := html.New("./templates", ".gohtml")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./public")

	storageRoute := app.Group("/api/storages")
	rootRoute := app.Group("/")

	inventoryService := services.NewInMemoryInventoryService()
	inventoryRoutes := routes.New(inventoryService)
	inventoryRoutes.DefineRoutes(storageRoute)
	inventoryRoutes.DefinePages(rootRoute)

	// Last middleware to match anything
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	return app.Listen(addr)
}

func NewFiberServer() Server {
	return &fiberServer{}
}
