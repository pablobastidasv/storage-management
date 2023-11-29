package server

import (
	"co.bastriguez/inventory/internal/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

type (
	Server struct {
		listenAddr string
		*handlers.StorageHandlers
	}
)

func NewFiberServer(listenAddr string, handler *handlers.StorageHandlers) *Server {
	return &Server{
		listenAddr:      listenAddr,
		StorageHandlers: handler,
	}
}

func (s *Server) Start() error {
	// Initialize standard Go html template engine
	engine := html.New("./templates", ".gohtml")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./public")

	app.Get("/", s.InventoryHomePageHandler)
	app.Get("/inventory/product/add-form", s.AddProductFormHandler)

	storageApi := app.Group("/api/storages")
	storageApi.Get("/main/products", s.GetProductsHandler)
	storageApi.Put("/main/products", s.PutProductsHandler)
	storageApi.Get("/main/remissions", s.StorageRemissionsHandler)

	// Last middleware to match anything
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	return app.Listen(s.listenAddr)
}
