package handlers

import (
	"time"
)

type (
	Product struct {
		Id           string
		Name         string
		Presentation string
	}

	ProductItem struct {
		Id     string
		Name   string
		Amount string
	}

	RemissionItem struct {
		ClientName string
		ProductItem
		CreatedAt time.Time
	}

	PutProductsRequest struct {
		Product string `json:"product"`
		Qty     int    `json:"qty"`
	}
)
