package handlers

import (
	"time"
)

const (
	hxTrigger             = "HX-Trigger"
	openRightDrawerEvent  = "open-right-drawer"
	closeRightDrawerEvent = "close-right-drawer"
)

type (
	Product struct {
		Id           string
		Name         string
		Presentation string
	}

	ProductItem struct {
		Id           string
		Name         string
		Amount       string
		Qty          int
		Presentation string
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
