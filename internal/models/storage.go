package models

import "time"

type Presentation int8
type RemissionState int8

const (
	KG Presentation = iota
	Grms
	Amount
)

const (
	Open RemissionState = iota
	Close
)

type (
	Storage struct {
		Id      string
		Content []InventoryContent
	}

	InventoryContent struct {
		Product Product
		Qty     int
	}

	Product struct {
		Id           string
		Name         string
		Presentation Presentation
	}

	Client struct {
		Id   string
		Name string
	}

	Remission struct {
		Id        string
		Product   Product
		Qty       int
		Client    Client
		State     RemissionState
		Return    int
		CreatedAt time.Time
		ClosedAt  time.Time
	}
)
