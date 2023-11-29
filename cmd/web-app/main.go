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
	db, err := databases.New()
	if err != nil {
		log.Fatalln(err)
	}

	// Persistence implemented interface
	storageRepo := repository.New(db)

	// Instance of a service
	storageService := services.NewStorage(storageRepo, "Nothing yet")

	// handlers instance
	storageHandler := handlers.NewStorage(storageService)

	// service routes
	app := server.NewFiberServer(addr, storageHandler)
	log.Fatal(app.Start())
}
