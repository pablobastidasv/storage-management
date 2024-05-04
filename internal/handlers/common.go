package handlers

import (
	"errors"

	"co.bastriguez/inventory/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

const (
)


func ErrorHandler(ctx *fiber.Ctx, err error) error {
	var wrongParameter *services.WrongParameter
	if errors.As(err, &wrongParameter) {
		ctx.Response().Header.Add(hxRetarget, errorAlertId)
		return ctx.Render("alert-messages", &AlertMessage{Message: err.Error(), Level: Danger})
	}

	log.Errorf("there was a unexpected error, it message is %s\n", err.Error())
	ctx.Status(200).Response().Header.Add("HX-Redirect", "https://http.cat/status/500")
	return nil
}

type AlertMessage struct {
	Message string
	Level   AlertMessageLevel
}
