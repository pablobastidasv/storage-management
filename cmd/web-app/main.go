package main

import (
	"co.bastriguez/inventory/internal/handlers"
	"log"
)

func main() {
	addr := ":8080"

	server := handlers.NewFiberServer()
	log.Fatal(server.Start(addr))
}
