package main

import (
	"co.bastriguez/inventory/internal"
	"co.bastriguez/inventory/internal/layouts"
)

func main() {
	addr := ":8080"

	layouts.Init("templates/layouts/", "templates/")
	internal.RunServer(addr)
}
