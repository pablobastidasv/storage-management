package components_test

import (
	"context"
	"fmt"
	"io"
	"testing"

	"co.bastriguez/inventory/internal/platform/server/views/components"
	"github.com/PuerkitoBio/goquery"
	"github.com/a-h/templ"
	"github.com/stretchr/testify/assert"
)

func Render(component templ.Component) (*goquery.Document, error) {
	r, w := io.Pipe()
	go func() {
		_ = component.Render(context.Background(), w)
		_ = w.Close()
	}()
	return goquery.NewDocumentFromReader(r)
}

func TestClientForm(t *testing.T) {
	doc, err := Render(components.ClientForm())
	assert.NoError(t, err, "Error when reading template")

	t.Run("client component is rendered with all fields", func(t *testing.T) {
		if doc.Find(`[data-testid="client-form"]`).Length() == 0 {
			t.Errorf("data-testid with value client-form is expected")
		}

		inputFields := []string{"doc_number", "name", "address", "email", "phone_number"}

		for _, n := range inputFields {
			selector := fmt.Sprintf(`input[name="%s"]`, n)
			if doc.Find(selector).Length() == 0 {
				t.Errorf("input with name '%s' is needed", n)
			}
		}

		if doc.Find(`select[name="doc_type"]`).Length() == 0 {
			t.Errorf("input with name 'id_type' is needed")
		}

	})

	t.Run("The form component does a hx-push to the /clients endpoint", func(t *testing.T) {
		postAddrs, exists := doc.Find(`[data-testid="client-form"]`).Attr("hx-post")
		assert.True(t, exists, "destination when post not defined")
		assert.Equal(t, "/clients", postAddrs, "form must post to /clients url")
	})
}
