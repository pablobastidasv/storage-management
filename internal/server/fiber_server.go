package server

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/html/v2"

	"co.bastriguez/inventory/internal/authenticator"
	"co.bastriguez/inventory/internal/handlers"
)

type (
	Server struct {
		listenAddr      string
		app             *fiber.App
		isAuthenticated func(c *fiber.Ctx) error
	}
)

func NewFiberServer(listenAddr string) *Server {
	// Initialize standard Go html template engine
	engine := html.New("./templates", ".go.html")

	app := fiber.New(fiber.Config{
		Views:        engine,
		ErrorHandler: handlers.ErrorHandler,
	})
	app.Use(logger.New())
	app.Use(recover.New())

	app.Static("/", "./public")

	server := &Server{
		listenAddr: listenAddr,
		app:        app,
	}

	server.loadAuth()

	return server
}

func (s *Server) loadAuth() {
	auth, err := authenticator.New()
	if err != nil {
		log.Fatalf("Error when loading authenticator %v.", err)
	}

	store := session.New()
	authHandlers := authenticator.NewHandlers(store, auth)
	s.isAuthenticated = authHandlers.IsAuthenticated

	s.app.Get("/login", authHandlers.HandleLogin)
	s.app.Get("/logout", authHandlers.HandleLogout)
	s.app.Get("/callback", authHandlers.HandleCallback)

	s.app.Get("/profile", s.isAuthenticated, authHandlers.HandleGetUsersMe)
}

func (s *Server) HandleProductsEndpoints(productHandler *handlers.ProductHandlers) {
	productsApi := s.app.Group("/products", s.isAuthenticated)

	productsApi.Get("/", productHandler.HandleGetProducts)
	productsApi.Post("/", productHandler.HandlePostProducts)
}

func (s *Server) HandleClientsEndpoints(clientHandlers *handlers.ClientHandlers) {
	clientsApi := s.app.Group("/clients", s.isAuthenticated)

	clientsApi.Post("/", clientHandlers.HandlePostClients)
}

func (s *Server) HandleStoragesEndpoints(storageHandler *handlers.StorageHandlers) {
	s.app.Get("/", s.isAuthenticated, storageHandler.HandleInventoryHomePage)

	storageApi := s.app.Group("/storages", s.isAuthenticated)
	storageApi.Get("/main/items/add", storageHandler.HandleAddProductFormFragment)
	storageApi.Get("/main/items", storageHandler.HandleGetItems)
	storageApi.Put("/main/items", storageHandler.HandlePutProducts)
	storageApi.Get(
		"/main/remissions",
		storageHandler.HandleGetRemissions,
	)
}

func (s *Server) HandleAdminEndpoints(adminHandler *handlers.AdminHandlers) {
	s.app.Get("/admin", s.isAuthenticated, adminHandler.HandleAdminHomePage)

	s.app.Get("/admin/products", s.isAuthenticated, adminHandler.HandleAdminProductsPage)
	s.app.Get("/admin/products/new", s.isAuthenticated, adminHandler.HandleAdminCreateProductFormFragment)

	s.app.Get("/admin/clients", s.isAuthenticated, adminHandler.HandleAdminClientsPage)
	s.app.Get("/admin/clients/new", s.isAuthenticated, adminHandler.HandleAdminCreateClientFormFragment)
}

func (s *Server) Start() error {
	// Last middleware to match anything
	s.app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	return s.app.Listen(s.listenAddr)
}

