package main

import (
	"co.bastriguez/inventory/internal/databases"
	"co.bastriguez/inventory/internal/repository"
	"co.bastriguez/inventory/internal/server"
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
	store, err := repository.New(db)
	if err != nil {
		log.Fatalln(err)
	}

	// Instance of a service

	// handlers instance

	// service routes
	app := server.NewFiberServer(addr)
	log.Fatal(app.Start())
}
