package models_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"co.bastriguez/inventory/internal/models"
)

func Test_ProductId_From(t *testing.T) {
	testCases := []struct {
		desc      string
		invalidId string
	}{
		{
			desc:      "when the ID is empty",
			invalidId: "",
		},
		{
			desc:      "when the ID is one space",
			invalidId: " ",
		},
		{
			desc: "when the ID is default empty string",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			_, err := models.ProductIdFrom(tC.invalidId)

			assert.NotNil(t, err)
			assert.Equal(
				t,
				"product id cannot be empty",
				err.Error(),
				"expected and Error() message must be equal",
			)
		})
	}
}

func Test_ProductIdFrom(t *testing.T) {
	testCases := []struct {
		desc    string
		validId string
	}{
		{
			desc:    "UUID format is valid",
			validId: "46406c0a-e7c4-44f3-812c-067ad2d4873f",
		},
		{
			desc:    "if the input has spaces at the begining those are trim",
			validId: " 46406c0a-e7c4-44f3-812c-067ad2d4873f",
		},
		{
			desc:    "if the input has spaces at the end those are trim",
			validId: "46406c0a-e7c4-44f3-812c-067ad2d4873f ",
		},
		{
			desc:    "if the input has spaces at the end and begining those are trim",
			validId: " 46406c0a-e7c4-44f3-812c-067ad2d4873f ",
		},
		{
			desc:    "if the input more than one space at the end or beginning, then those are trimmed",
			validId: "      46406c0a-e7c4-44f3-812c-067ad2d4873f    ",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result, err := models.ProductIdFrom(tC.validId)

			assert.Nil(t, err, "the error must be nil")
			assert.NotNil(t, result, "the resulting information must not be nil")
			assert.Equal(
				t,
				"46406c0a-e7c4-44f3-812c-067ad2d4873f",
				result.ToString(),
				"Id value is the expected without spaces",
			)
		})
	}
}
