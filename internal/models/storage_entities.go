package models

import (
	"fmt"
	"time"
)

type (
	Presentation   string
	RemissionState int8

	ProductName string
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
		Id           ProductId
		Name         ProductName
		Presentation Presentation
	}

	Transaction struct {
		Datetime time.Time
		Product  Product
		Storage  Storage
		Qty      int
	}
)

func NewProduct(id string, name string, presentation string) (*Product, error) {
	valId, err := ProductIdFrom(id)
	if err != nil {
		return nil, err
	}

	valPresentation, err := NewPresentation(presentation)
	if err != nil {
		return nil, err
	}

	valName, err := ProductNameFrom(name)
	if err != nil {
		return nil, err
	}

	return &Product{
		Id:           valId,
		Name:         valName,
		Presentation: valPresentation,
	}, nil
}

func CreateProduct(id string, name string, presentation string) (*Product, error) {
	return NewProduct(id, name, presentation)
}

func ProductNameFrom(name string) (ProductName, error) {
	if len(name) == 0 {
		return ProductName(""), NewDomainError("Product name cannot be empty")
	}
	return ProductName(name), nil
}

func NewPresentation(presentation string) (Presentation, error) {
	switch presentation {
	case "KG":
		return KG, nil
	case "GRAMS":
		return Grms, nil
	case "QUANTITY":
		return Amount, nil
	default:
		return unknown, NewDomainError(fmt.Sprintf("'%s' is an invalid presentation", presentation))
	}
}

func ListPresentations() []Presentation {
	return []Presentation{
		KG, Grms, Amount,
	}
}

func (p *Presentation) ToString() string {
	return string(*p)
}

func (n *ProductName) ToString() string {
	return string(*n)
}
