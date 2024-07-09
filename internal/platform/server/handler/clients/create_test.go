package clients_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"co.bastriguez/inventory/internal/creating"
	"co.bastriguez/inventory/internal/inventory"
	"co.bastriguez/inventory/internal/platform/server/handler/clients"
	"co.bastriguez/inventory/kit/command/commandmocks"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestPostClientHandler_Create(t *testing.T) {
	bus := new(commandmocks.Bus)

	app := fiber.New()

	app.Post("/clients", clients.PostClientHandler(bus))

	t.Run("given invalid request it returns 422", func(t *testing.T) {
		bus.ExpectedCalls = nil
		bus.On("DispatchCommand", mock.Anything, mock.Anything).Return(inventory.ErrInvalidClientName)

		data := "id_type=CC&id_number=1234567890"

		req := httptest.NewRequest("POST", "/clients", strings.NewReader(data))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		response, err := app.Test(req)
		require.Nil(t, err, "Error when testing request.")

		assert.Equal(t, http.StatusUnprocessableEntity, response.StatusCode)
	})

	t.Run("given a valid request it returns 201", func(t *testing.T) {
		bus.ExpectedCalls = nil

		expectedCommand := creating.NewCreateClientCommand("CC", "1234567890", "Jose")
		bus.On("DispatchCommand", mock.Anything, expectedCommand).Return(nil)

		data := "id_type=CC&id_number=1234567890&name=Jose"

		req := httptest.NewRequest("POST", "/clients", strings.NewReader(data))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		response, err := app.Test(req)
		require.Nil(t, err, "Error when testing request.")

		assert.Equal(t, http.StatusCreated, response.StatusCode)
		bus.AssertExpectations(t)
	})

}
