package models

import (
	"fmt"
	"time"
)

type (
	Presentation   string
	RemissionState int8
)

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

func CreateProduct(id string, name string, presentation string) (*Product, error) {
	if NewPresentation(presentation) == unknown {
		return nil, &DomainError{
			desc: fmt.Sprintf("'%s' is a invalid presentation", presentation),
		}
	}

	if len(name) == 0 {
		return nil, &DomainError{
			desc: "Product name cannot be empty",
		}
	}

	return &Product{
		Id:           id,
		Name:         name,
		Presentation: Presentation(presentation),
	}, nil
}

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

func ListPresentations() []Presentation {
	return []Presentation{
		KG, Grms, Amount,
	}
}

type DomainError struct {
	desc string
}

func (e *DomainError) Error() string {
	return e.desc
}
