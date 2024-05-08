package handlers

import (
	"time"
)

const (
	hxRetarget   = "HX-Retarget"
	hxTrigger    = "HX-Trigger"
	errorAlertId = "#error-alert"
	messagesId   = "#validation-messages"

	openRightDrawerEvent  = "open-right-drawer"
	closeRightDrawerEvent = "close-right-drawer"

	Primary   AlertMessageLevel = "primary"
	Secondary                   = "secondary"
	Success                     = "success"
	Danger                      = "danger"
	Warning                     = "warning"
	Info                        = "info"
	Light                       = "light"
	Dark                        = "dark"
)

type (
	AlertMessageLevel string

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
