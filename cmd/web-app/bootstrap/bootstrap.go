package bootstrap

import (
	"fmt"
	"log/slog"
	"os"

	"co.bastriguez/inventory/internal/creating"
	"co.bastriguez/inventory/internal/platform/bus/inmemory"
	"co.bastriguez/inventory/internal/platform/server"
	"co.bastriguez/inventory/internal/platform/storage/postgres"
	"github.com/joho/godotenv"
)

const (
	EnvVar = "BASTRIGUEZ_ENV"
)

func Run() error {
	return withLog(func() error {
		loadEnv()
		return start()
	})
}

func start() error {
	db, err := dbOpen()
	if err != nil {
		return err
	}

	clientRepository := postgres.NewClientRepository(db) // TODO: time.Duration for timeout

	bus := inmemory.New()
	clientService := creating.NewClientService(clientRepository)

	createClientCommandHandler := creating.NewCreateClientCommandHandler(clientService)
	bus.RegisterCommandHandler(creating.CreateCommandType, createClientCommandHandler)

	srv := server.New(host, port, clientRepository, bus)
	return srv.Run()
}

func loadEnv() {
	env := os.Getenv(EnvVar)
	if "" == env {
		env = "dev"
	}

	err := godotenv.Load(fmt.Sprintf(".env.%s.local", env))
	if err != nil {
		slog.Debug("error loading file env file: ", slog.Any("error", err))
	}

	if "test" != env {
		err = godotenv.Load(".env.local")
		if err != nil {
			slog.Debug("error loading file env file: ", slog.Any("error", err))
		}
	}
	err = godotenv.Load(fmt.Sprintf(".env.%s", env))
	if err != nil {
		slog.Debug("error loading file env file: ", slog.Any("error", err))
	}

	err = godotenv.Load() // The original .env
	if err != nil {
		slog.Debug("error loading file env file: ", slog.Any("error", err))
	}
}
