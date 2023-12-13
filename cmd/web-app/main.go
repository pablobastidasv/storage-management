package main

import (
	"co.bastriguez/inventory/internal/databases"
	"co.bastriguez/inventory/internal/handlers"
	"co.bastriguez/inventory/internal/repository"
	"co.bastriguez/inventory/internal/server"
	"co.bastriguez/inventory/internal/services"
	"log"
)

func main() {
	addr := ":8080"

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
