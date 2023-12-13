package server

import (
	"co.bastriguez/inventory/internal/handlers"
	"co.bastriguez/inventory/internal/services"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
	"log"
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
		Views:        engine,
		ErrorHandler: ErrorHandler,
	})
	app.Use(logger.New())

	app.Static("/", "./public")

	return &Server{
		listenAddr: listenAddr,
		app:        app,
	}
}

func (s *Server) HandleProductsEndpoints(productHandler *handlers.ProductHandlers) {
	productsApi := s.app.Group("/api/products")

	productsApi.Get("/", productHandler.HandleGetProducts)
}

func (s *Server) HandleStoragesEndpoints(storageHandler *handlers.StorageHandlers) {
	s.app.Get("/", storageHandler.HandleInventoryHomePage)
	s.app.Get("/inventory/product/add-form", storageHandler.HandleAddProductFormFragment)

	storageApi := s.app.Group("/api/storages")
	storageApi.Get("/main/products", storageHandler.HandleGetProducts)
	storageApi.Put("/main/products", storageHandler.HandlePutProducts)
	storageApi.Get("/main/remissions", storageHandler.HandleGetRemissions) // TODO: move to remissions
}

func (s *Server) HandleAdminEndpoints(adminHandler *handlers.AdminHandlers) {
	s.app.Get("/admin", adminHandler.HandleAdminHomePage)
	s.app.Get("/admin/product/create-form", adminHandler.HandleAdminCreateProductFormFragment)
}

func (s *Server) Start() error {
	// Last middleware to match anything
	s.app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	return s.app.Listen(s.listenAddr)
}

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	var wrongParameter *services.WrongParameter
	if errors.As(err, &wrongParameter) {
		ctx.Response().Header.Add("HX-Retarget", "#error-alert")
		return ctx.Render("alert-messages", &AlertMessage{Message: err.Error()})
	}

	log.Printf("there was a unexpected error, it message is %s", err.Error())
	ctx.Status(200).Response().Header.Add("HX-Redirect", "https://http.cat/status/500")
	return nil
}

type AlertMessage struct {
	Message string
}
