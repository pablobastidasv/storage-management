package main

import (
	"log"

	"co.bastriguez/inventory/cmd/web-app/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatalf("error starting the server: %v", err)
	}
}
