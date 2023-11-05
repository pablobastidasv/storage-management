package handlers

import (
	"log"
	"net/http"

	"co.bastriguez/inventory/internal/components"
	"co.bastriguez/inventory/internal/services"
	"github.com/a-h/templ"
	"github.com/gorilla/mux"
)

func Start(addr string) {
	r := mux.NewRouter()

	r.Handle("/", templ.Handler(components.Index()))

	service := services.NewInMemoryInventoryService()
	inventoryHandler := NewInventoryHandler(service)

	r.HandleFunc("/inventory", inventoryHandler.Index)

	http.Handle("/", r)

	// serving static files
	fs := http.FileServer(http.Dir("statics"))
	http.Handle("/statics/", http.StripPrefix("/statics/", fs))

	log.Println("Starting up on 8080")
	log.Fatal(http.ListenAndServe(addr, nil))
}
