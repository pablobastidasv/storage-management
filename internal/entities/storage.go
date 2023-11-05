package entities

type Presentation int8

const (
	KG Presentation = iota
	Grms
	Amount
)

type (
	Storage struct {
		Id      string
		Content []InventoryContent
	}

	InventoryContent struct {
		Product Product
		Qty     int16
	}

	Product struct {
		Id           string
		Name         string
		Presentation Presentation
	}
)
