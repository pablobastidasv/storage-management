package main

import (
	"co.bastriguez/inventory/internal/server"
	"log"
)

func main() {
	addr := ":8080"

	server := server.NewFiberServer()
	log.Fatal(server.Start(addr))
}
