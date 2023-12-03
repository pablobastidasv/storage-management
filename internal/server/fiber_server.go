package server

import (
	"co.bastriguez/inventory/internal/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
)

type (
	Server struct {
		listenAddr string
		app        *fiber.App
	}
)

func NewFiberServer(listenAddr string) *Server {
	// Initialize standard Go html template engine
	engine := html.New("./templates", ".gohtml")

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Use(logger.New())

	app.Static("/", "./public")

	return &Server{
		listenAddr: listenAddr,
		app:        app,
	}
}

func (s *Server) ProductHandler(productHandler *handlers.ProductHandlers) {
	productsApi := s.app.Group("/api/products")

	productsApi.Get("/", productHandler.HandleGetProducts)
}

func (s *Server) StorageHandler(storageHandler *handlers.StorageHandlers) {
	s.app.Get("/", storageHandler.InventoryHomePageHandler)
	s.app.Get("/inventory/product/add-form", storageHandler.AddProductFormHandler)

	storageApi := s.app.Group("/api/storages")
	storageApi.Get("/main/products", storageHandler.GetProductsHandler)
	storageApi.Put("/main/products", storageHandler.HandlePutProducts)
	storageApi.Get("/main/remissions", storageHandler.StorageRemissionsHandler)

}

func (s *Server) Start() error {
	// Last middleware to match anything
	s.app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	return s.app.Listen(s.listenAddr)
}
