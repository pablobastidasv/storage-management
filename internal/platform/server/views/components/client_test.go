package components_test

import (
	"context"
	"io"
	"testing"

	"co.bastriguez/inventory/internal/platform/server/views/components"
	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"
)

func TestClientForm(t *testing.T) {

	t.Run("client component is rendered", func(t *testing.T) {
		r, w := io.Pipe()
		go func() {
			_ = components.ClientForm().Render(context.Background(), w)
			_ = w.Close()
		}()
		doc, err := goquery.NewDocumentFromReader(r)
		assert.NoError(t, err, "Error when reading template")

		if doc.Find(`[data-testid="client-form"]`).Length() == 0 {
			t.Errorf("data-testid with value client-form is expected")
		}
	})

	t.Run("the document type field is a select and it's named 'id_type'", func(t *testing.T) {
		r, w := io.Pipe()
		go func() {
			_ = components.ClientForm().Render(context.Background(), w)
			_ = w.Close()
		}()
		doc, err := goquery.NewDocumentFromReader(r)
		assert.NoError(t, err, "Error when reading template")

		if doc.Find(`select[name="id_type"]`).Length() == 0 {
			t.Errorf("input with name id_type is needed")
		}
	})

	t.Run("the document number field is name 'id_number'", func(t *testing.T) {
		r, w := io.Pipe()
		go func() {
			_ = components.ClientForm().Render(context.Background(), w)
			_ = w.Close()
		}()
		doc, err := goquery.NewDocumentFromReader(r)
		assert.NoError(t, err, "Error when reading template")

		if doc.Find(`input[name="id_number"]`).Length() == 0 {
			t.Errorf("input with name id_number is needed")
		}
	})

    t.Run("the name field is named 'name'", func(t *testing.T) {
		r, w := io.Pipe()
		go func() {
			_ = components.ClientForm().Render(context.Background(), w)
			_ = w.Close()
		}()
		doc, err := goquery.NewDocumentFromReader(r)
		assert.NoError(t, err, "Error when reading template")

		if doc.Find(`input[name="name"]`).Length() == 0 {
			t.Errorf("input with name 'name' is needed")
		}
    })
}
