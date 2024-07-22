package components_test

import (
	"testing"

	"co.bastriguez/inventory/internal/platform/server/views/components"
	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"
)

func TestInput(t *testing.T) {
	name := "inputName"

	attrs := components.InputAttrs{
		Label: "Name:",
	}

	doc, err := Render(components.Input(name, attrs))
	assert.NoError(t, err, "Error rendering component")

	label := doc.Find(`[data-testid="label"]`)
	assert.Equal(t, 1, label.Length(), "Only one label must be present")

	input := doc.Find(`[data-testid="input"]`)
	assert.Equal(t, 1, input.Length(), "only one input must be present")

	t.Run("label contains the text of the label", func(t *testing.T) {
		assert.Equal(t, attrs.Label, label.Text(), "Label text must be the given one")
	})

	t.Run("label for value is the given input name", func(t *testing.T) {
		attr, exists := label.Attr("for")
		assert.True(t, exists, "the for attribute should exists")
		assert.Equal(t, name, attr, "for value should be the name of the fields")
	})

	t.Run("input name is the given value name", func(t *testing.T) {
		attr, exists := input.Attr("name")
		assert.True(t, exists, "the attribute attribute should exists")
		assert.Equal(t, name, attr, "name input value should be the name of the fields")
	})
}

func TestSelect(t *testing.T) {
	name := "selectName"
	attrs := components.SelectAttrs{
		Label: "Select something:",
		Options: []components.SelectOption{
			{
				Value: "a",
				Label: "Letter A",
			},
			{
				Value: "b",
				Label: "Letter B",
			},
		},
	}

	doc, err := Render(components.Select(name, attrs))
	assert.NoError(t, err, "Error rendering component")

	label := doc.Find(`[data-testid="label"]`)
	assert.Equal(t, 1, label.Length(), "only one label is expected")

	sel := doc.Find(`[data-testid="select"]`)
	assert.Equal(t, 1, sel.Length(), "only one select is expected")

	t.Run("label contain the given dalue in attributes", func(t *testing.T) {
		assert.Equal(t, "Select something:", label.Text(), "label should have Label attribute")
	})

	t.Run("label must have the name in its for attribute", func(t *testing.T) {
		forValue, exists := label.Attr("for")
		assert.True(t, exists, "for attribute should be in the label")
		assert.Equal(t, name, forValue, "for value should be the given name")
	})

	t.Run("select name is the given value", func(t *testing.T) {
		nameValue, exists := sel.Attr("name")
		assert.True(t, exists, "name attribute should be in the select")
		assert.Equal(t, name, nameValue, "name attribute value should be the given value")
	})

	t.Run("select's options ammount is 2", func(t *testing.T) {
		opts := doc.Find("option")
		assert.Equal(t, 2, opts.Length())
	})

	t.Run("select's options labels has the expected values", func(t *testing.T) {
		opts := doc.Find("option")

		var labels []string
		for _, n := range opts.Nodes {
			labels = append(labels, n.FirstChild.Data)
		}
		assert.Equal(t, []string{"Letter A", "Letter B"}, labels, "labels should be the given label value")
	})

	t.Run("select options' values are the given one", func(t *testing.T) {
		opts := doc.Find("option")

		var values []string
		opts.Each(func(_ int, s *goquery.Selection) {
			v, e := s.Attr("value")
			assert.True(t, e, "all option elements should have value")
			values = append(values, v)
		})
		assert.Equal(t, []string{"a", "b"}, values, "option values should be the given value")
	})
}

func TestButton(t *testing.T) {

	attrs := components.ButtonAttrs{
		Label: "Create",
	}
	doc, err := Render(components.Button(attrs))
	assert.NoError(t, err, "component should be rendered")

	button := doc.Find(`[data-testid="button"]`)
	assert.Equal(t, 1, button.Length(), "there most be an identifiable element dalled button")

	tests := []struct {
		Name   string
		Attr   components.ButtonAttrs
		Assert func(*goquery.Selection, *testing.T)
	}{
		{
			Name: "given a label, it should be in the component",
			Attr: components.ButtonAttrs{
				Label: "Create",
			},
			Assert: func(button *goquery.Selection, tt *testing.T) {
				assert.Equal(tt, "Create", button.Text(), "button text should be 'Create'")
			},
		},
		{
			Name: "when type is provided, type is present in the component",
			Attr: components.ButtonAttrs{
				Label: "Something",
				Type:  "submit",
			},
			Assert: func(button *goquery.Selection, tt *testing.T) {
				value, exists := button.Attr("type")
				assert.True(tt, exists, "type attribute should be present")
				assert.Equal(tt, "submit", value, "submit value should be the value of the type attribute")
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			doc, err := Render(components.Button(test.Attr))
			assert.NoError(t, err, "component should be rendered")

			button := doc.Find(`[data-testid="button"]`)
			test.Assert(button, t)
		})
	}

}
