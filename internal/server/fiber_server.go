package server

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"

	"co.bastriguez/inventory/internal/handlers"
	"co.bastriguez/inventory/internal/services"
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
	productsApi := s.app.Group("/products")

	productsApi.Get("/", productHandler.HandleGetProducts)
	productsApi.Post("/", productHandler.HandlePostProducts)
}

func (s *Server) HandleStoragesEndpoints(storageHandler *handlers.StorageHandlers) {
	s.app.Get("/", storageHandler.HandleInventoryHomePage)

	storageApi := s.app.Group("/storages")
	storageApi.Get("/main/items/add", storageHandler.HandleAddProductFormFragment)
	storageApi.Get("/main/items", storageHandler.HandleGetProducts)
	storageApi.Put("/main/items", storageHandler.HandlePutProducts)
	storageApi.Get(
		"/main/remissions",
		storageHandler.HandleGetRemissions,
	)
}

func (s *Server) HandleAdminEndpoints(adminHandler *handlers.AdminHandlers) {
	s.app.Get("/admin", adminHandler.HandleAdminHomePage)
	s.app.Get("/products/new", adminHandler.HandleAdminCreateProductFormFragment)
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
