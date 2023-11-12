package handlers

import (
	"co.bastriguez/inventory/internal/components"
	"co.bastriguez/inventory/internal/server"
	"co.bastriguez/inventory/internal/services"
	"github.com/a-h/templ"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type muxServer struct{}

func NewMuxServer() server.Server {
	return &muxServer{}
}

func (t muxServer) Start(addr string) error {
	r := mux.NewRouter()

	r.Handle("/", templ.Handler(components.Index()))

	service := services.NewInMemoryInventoryService()
	inventoryHandler := NewInventoryHandler(service)

	r.HandleFunc("/inventory", inventoryHandler.Index)
	r.HandleFunc("/inventory/create-remission", inventoryHandler.CreateRemission)

	http.Handle("/", r)

	// serving static files
	fs := http.FileServer(http.Dir("components"))
	http.Handle("/statics/", http.StripPrefix("/components/", fs))

	log.Println("Starting up on 8080")

	return http.ListenAndServe(addr, nil)
}
