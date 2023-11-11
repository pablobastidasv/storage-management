package handlers

import (
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

	app.Get("/hello", RootHandler)

	app.Get("/storages/main/products", StorageProductsHandler)
	app.Get("/storages/main/remissions", StorageRemissionsHandler)

	// Last middleware to match anything
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	return app.Listen(addr)
}

func StorageRemissionsHandler(ctx *fiber.Ctx) error {
	return ctx.Render("remissions", nil)
}

func StorageProductsHandler(ctx *fiber.Ctx) error {
	return ctx.Render("inventory", nil)
}

func RootHandler(ctx *fiber.Ctx) error {
	return ctx.SendString("Hello world!!!")
}

func NewFiberServer() Server {
	return &fiberServer{}
}
