package main

import (
	"co.bastriguez/inventory/internal/server"
	"log"
)

func main() {
	addr := ":8080"

	app := server.NewFiberServer(addr)
	log.Fatal(app.Start())
}
