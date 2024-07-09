package server

import (
	"fmt"
	"log"
	"log/slog"

	"co.bastriguez/inventory/internal/inventory"
	"co.bastriguez/inventory/internal/platform/server/handler/clients"
	"co.bastriguez/inventory/kit/command"
	"github.com/gofiber/fiber/v2"
	slogfiber "github.com/samber/slog-fiber"
)

type Server struct {
	httpAddrs string
	engine    *fiber.App

	bus command.Bus

	// dependencies
	inventory.ClientRepository
}

func New(host string, port uint, repository inventory.ClientRepository, bus command.Bus) Server {
	engine := fiber.New()
    engine.Use(slogfiber.New(slog.Default()))

	srv := Server{
		engine:    engine,
		httpAddrs: fmt.Sprintf("%s:%d", host, port),
		bus:       bus,
	}

	srv.registerRoutes()
	return srv
}

func (s *Server) Run() error {
	log.Println("Server runnin on", s.httpAddrs)
	return s.engine.Listen(s.httpAddrs)
}

func (s *Server) registerRoutes() {
	s.engine.Post("/clients", clients.PostClientHandler(s.bus))
}
