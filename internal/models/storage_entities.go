package models

import "time"

type Presentation string
type RemissionState int8

const (
	KG      Presentation = "KG"
	Grms    Presentation = "GRAMS"
	Amount  Presentation = "QTY"
	unknown Presentation = "UNKNOWN"
)

type (
	Storage struct {
		Id      string
		Content []InventoryItem
	}

	InventoryItem struct {
		Product InventoryProduct
		Qty     int
	}

	InventoryProduct struct {
		Id           string
		Name         string
		Presentation Presentation
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
