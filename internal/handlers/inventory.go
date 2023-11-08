package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"co.bastriguez/inventory/internal/components"
	"co.bastriguez/inventory/internal/entities"
	"co.bastriguez/inventory/internal/services"
)

type (
	WebInventoryHandler struct {
		service services.InventoryService
	}

	InventoryHandler interface {
		Index(w http.ResponseWriter, r *http.Request)
	}
)

func NewInventoryHandler(inventoryService services.InventoryService) InventoryHandler {
	return &WebInventoryHandler{
		service: inventoryService,
	}
}

func (h WebInventoryHandler) Index(w http.ResponseWriter, r *http.Request) {
	var items []components.Product
	for _, content := range h.service.RetrieveInventory() {
		ammount := strconv.FormatInt(int64(content.Qty), 10)
		items = append(items, components.Product{
			Id:     content.Product.Id,
			Name:   content.Product.Name,
			Amount: ammount,
			Unit:   translateUnit(content.Product.Presentation),
		})
	}

	var remissions []components.Remission

	for _, remission := range h.service.RetrieveOpenRemissions() {
		ammount := fmt.Sprintf("%d %s", remission.Qty, translateUnit(remission.Product.Presentation))
		remissions = append(remissions, components.Remission{
			Id:         remission.Id,
			ClientName: remission.Client.Name,
			Amount:     ammount,
			Product: components.Product{
				Id:   remission.Product.Id,
				Name: remission.Product.Name,
			},
		})
	}

	inventory := components.Inventory{
		Products:   items,
		Remissions: remissions,
	}
	components.InventoryMain(inventory).Render(r.Context(), w)
}

func translateUnit(unit entities.Presentation) string {
	switch unit {
	case entities.KG:
		return "Kilogramos"
	case entities.Amount:
		return "Cantidad"
	case entities.Grms:
		return "Gramos"
	default:
		return "Desconocido"
	}
}
