package server

import (
	"co.bastriguez/inventory/internal/handlers"
	"co.bastriguez/inventory/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

type (
	Server struct {
		listenAddr string
	}
)

func NewFiberServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
	}
}

func (s *Server) Start() error {
	inventoryService := services.NewInMemoryService()
	storageHandler := handlers.NewStorageHandler(inventoryService)

	// Initialize standard Go html template engine
	engine := html.New("./templates", ".gohtml")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./public")

	app.Get("/", storageHandler.InventoryHomePageHandler)
	app.Get("/inventory/product/add-form", storageHandler.AddProductFormHandler)

	storageApi := app.Group("/api/storages")
	storageApi.Get("/main/products", storageHandler.GetProductsHandler)
	storageApi.Put("/main/products", storageHandler.PutProductsHandler)
	storageApi.Get("/main/remissions", storageHandler.StorageRemissionsHandler)

	// Last middleware to match anything
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	return app.Listen(s.listenAddr)
}
