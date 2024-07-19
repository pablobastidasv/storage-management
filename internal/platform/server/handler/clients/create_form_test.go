package clients_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"co.bastriguez/inventory/internal/platform/server/handler/clients"
	"github.com/PuerkitoBio/goquery"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetClientFormHandler_Create(t *testing.T) {
	app := fiber.New()

	app.Get("/clients/new", clients.GetClientFormHandler())

	t.Run("given a request, then it returns http code 200 and contain the client form", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/clients/new", nil)

		response, err := app.Test(req)
		require.NoError(t, err, "Error when testing request.")

		assert.Equal(t, http.StatusOK, response.StatusCode)

		doc, err := goquery.NewDocumentFromReader(response.Body)
		assert.NoError(t, err, "Error when reading the response payload.")

		if doc.Find(`[data-testid="client-form"]`).Length() == 0 {
			t.Error("expected client form to be rendered, but it wasn't")
		}
	})

}
