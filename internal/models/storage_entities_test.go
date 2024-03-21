package models_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"co.bastriguez/inventory/internal/models"
)

func Test_CreateProduct_OK(t *testing.T) {
	type inputValue struct {
		id           string
		name         string
		presentation string
	}

	testCases := []struct {
		input inputValue
	}{
		{
			input: inputValue{
				id:           "A",
				name:         "Coper",
				presentation: "GRAMS",
			},
		},
	}
	for _, tC := range testCases {
		t.Run("Given a valid information, then the product is created", func(t *testing.T) {
			product, err := models.CreateProduct(
				tC.input.id,
				tC.input.name,
				tC.input.presentation,
			)

			assert.Nil(t, err, "must not give any error")
			assert.Equal(t, tC.input.id, product.Id.ToString(), "id must be the given")
			assert.Equal(t, models.ProductName(tC.input.name), product.Name, "name must be the given")
			assert.Equal(
				t,
				tC.input.presentation,
				product.Presentation.ToString(),
				"presentation must be the given",
			)
		})
	}
}
