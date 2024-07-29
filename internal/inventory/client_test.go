package inventory_test

import (
	"testing"

	"co.bastriguez/inventory/internal/inventory"
	"github.com/stretchr/testify/assert"
)

func Test_ClientEmail_ShouldNotBeBlank(t *testing.T) {
	_, err := inventory.NewClientEmail("")

	assert.Error(t, err)

	assert.ErrorIs(t, err, inventory.ErrEmptyEmail)
	assert.ErrorIs(t, err, inventory.ErrInvalidEmail)
    assert.ErrorAs(t, err, inventory.ValidationError)
}
