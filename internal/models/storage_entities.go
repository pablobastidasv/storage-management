package models

import "time"

type Presentation int8
type RemissionState int8

const (
	KG Presentation = iota
	Grms
	Amount
	unknown
)

type (
	Storage struct {
		Id      string
		Content []InventoryItem
	}

	InventoryItem struct {
		Product Product
		Qty     int
	}

	Product struct {
		Id           string
		Name         string
		Presentation Presentation
	}

	Transaction struct {
		Datetime time.Time
		Product  Product
		Storage  Storage
		Qty      int
	}
)

func NewPresentation(presentation string) Presentation {
	switch presentation {
	case "KG":
		return KG
	case "GRAMS":
		return Grms
	case "QUANTITY":
		return Amount
	default:
		return unknown
	}
}
