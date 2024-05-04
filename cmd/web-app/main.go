package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"co.bastriguez/inventory/internal/databases"
	"co.bastriguez/inventory/internal/handlers"
	"co.bastriguez/inventory/internal/repository"
	"co.bastriguez/inventory/internal/server"
	"co.bastriguez/inventory/internal/services"
)

func main() {
	loadEnv()

	addr := os.Getenv("ADDRS")

	// database access
	db, err := databases.NewMongo()
	if err != nil {
		log.Fatalln(err)
	}

	fiberServer := server.NewFiberServer(addr)

	// Product service
	productRepo := repository.NewMongoProductsRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	fiberServer.HandleProductsEndpoints(productHandler)

    // Client service
    clientRepository := repository.NewClientMongoRepository(db)
    clientService := services.NewClientService(clientRepository)
    clientHandler := handlers.NewClientHandler(clientService)

    fiberServer.HandleClientsEndpoints(clientHandler)

	// Persistence implemented interface
	storageRepo := repository.NewStorageMongoRepository(db)
	storageService := services.NewStorageService(storageRepo, productRepo)
	storageHandler := handlers.NewStorageHandler(storageService, productService)

	fiberServer.HandleStoragesEndpoints(storageHandler)

	// Admin pages
	adminHandlers := handlers.NewAdminHandler()

	fiberServer.HandleAdminEndpoints(adminHandlers)

	// start service
	log.Fatal(fiberServer.Start())
}

func loadEnv() {
	env := os.Getenv("BASTRIGUEZ_ENV")
	if "" == env {
		env = "dev"
	}

	err := godotenv.Load(fmt.Sprintf(".env.%s.local", env))
	if err != nil {
		log.Printf("Error loading file env file: %v\n", err)
	}

	if "test" != env {
		err = godotenv.Load(".env.local")
		if err != nil {
			log.Printf("Error loading file env file: %v\n", err)
		}
	}
	err = godotenv.Load(fmt.Sprintf(".env.%s", env))
	if err != nil {
		log.Printf("Error loading file env file: %v\n", err)
	}

	err = godotenv.Load() // The original .env
	if err != nil {
		log.Printf("Error loading file env file: %v\n", err)
	}
}
