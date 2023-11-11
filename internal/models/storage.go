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
		Qty     int16
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
		Qty       int16
		Client    Client
		State     RemissionState
		Return    int16
		CreatedAt time.Time
		ClosedAt  time.Time
	}
)
