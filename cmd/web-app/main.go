package main

import "co.bastriguez/inventory/internal/handlers"

func main() {
	addr := ":8080"

	handlers.Start(addr)
}
