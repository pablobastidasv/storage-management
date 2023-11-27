package main

import (
	"co.bastriguez/inventory/internal/databases"
	"co.bastriguez/inventory/internal/server"
	"co.bastriguez/inventory/internal/storage"
	"log"
)

func main() {
	addr := ":8080"

	db, err := databases.New()
	if err != nil {
		log.Fatalln(err)
	}

	store, err := storage.New(db)
	if err != nil {
		log.Fatalln(err)
	}

	app := server.NewFiberServer(addr)
	log.Fatal(app.Start())
}
