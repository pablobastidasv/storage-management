package inventory_test

import (
	"fmt"
	"testing"

	"co.bastriguez/inventory/internal/inventory"
	"github.com/stretchr/testify/assert"
)

func Test_ClientName_ShouldNotBeBlank(t *testing.T) {
	tests := []struct {
		name    string
		action  func() error
		field   string
		message string
	}{
		{
			name: "Given an empty name, Then field is name and message is related to it",
			action: func() error {
				_, err := inventory.NewClientName("")
				return err
			},
			field:   "name",
			message: "no puede estar vacio",
		},
		{
			name: "Given only spaces on the name, then error is returned",
			action: func() error {
				_, err := inventory.NewClientName("  ")
				return err
			},
			field:   "name",
			message: "no puede estar vacio",
		},
		// Email validations
		{
			name: "Given an empty email, then field is email and message is related to it",
			action: func() error {
				_, err := inventory.NewClientEmail("")
				return err
			},
			field:   "email",
			message: "no puede estar vacio",
		},
		{
			name: "Given only spaces on the email, then error is returned",
			action: func() error {
				_, err := inventory.NewClientEmail("  ")
				return err
			},
			field:   "email",
			message: "no puede estar vacio",
		},
		{
			name: "Given only spaces on the name, then error is returned",
			action: func() error {
				_, err := inventory.NewClientEmail("invalidemail")
				return err
			},
			field:   "email",
			message: "formato de email no es valido",
		},
		// Document type
		{
			name: "Given an empty document type, then field is doc_type and message is related to it",
			action: func() error {
				_, err := inventory.NewDocumentType("")
				return err
			},
			field:   "doc_type",
			message: "no puede estar vacio",
		},
		{
			name: "Given only spaces on the document type, then error is returned",
			action: func() error {
				_, err := inventory.NewDocumentType("  ")
				return err
			},
			field:   "doc_type",
			message: "no puede estar vacio",
		},
		{
			name: "Given a invalid document type, then error is returned",
			action: func() error {
				_, err := inventory.NewDocumentType("BSN")
				return err
			},
			field:   "doc_type",
			message: "document type 'BSN' is not valid",
		},
		// Documennt number
		{
			name: "Given an empty document number, then field is doc_number and message is related to it",
			action: func() error {
				_, err := inventory.NewDocumentNumber("")
				return err
			},
			field:   "doc_number",
			message: "no puede estar vacio",
		},
		{
			name: "Given only spaces on the document number, then error is returned",
			action: func() error {
				_, err := inventory.NewDocumentNumber("  ")
				return err
			},
			field:   "doc_number",
			message: "no puede estar vacio",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			err := test.action()

			if err == nil {
				tt.Fatal("error was expected to be returned")
			}

			result := new(inventory.FieldError)
			assert.ErrorAs(tt, err, &result, "error must be a FieldError")

			assert.Equal(tt, test.field, result.Field)
			assert.Equal(tt, test.message, result.Messsage)
		})

	}
}

func TestDocNumber_InvalidDocumentNumber(t *testing.T) {
	testCases := []struct {
		invalidValue string
	}{
		{
			invalidValue: "invalid",
		},
		{
			invalidValue: "1130468452-25",
		},
		{
			invalidValue: "0321654",
		},
		{
			invalidValue: "12",
		},
	}
	for _, tc := range testCases {
		desc := fmt.Sprintf("invalid document number '%s'", tc.invalidValue)
		t.Run(desc, func(tt *testing.T) {
			_, err := inventory.NewDocumentNumber(tc.invalidValue)

			if err == nil {
				t.Fatal("error is expected")
			}

			result := new(inventory.FieldError)
			assert.ErrorAs(tt, err, &result, "error must be a Field Error")

			expectedErrorMessage := fmt.Sprintf("numero de documento '%s' no es valido", tc.invalidValue)
			assert.Equal(tt, "doc_number", result.Field)
            assert.Equal(tt, expectedErrorMessage, result.Messsage)
		})
	}
}

func TestDocNumber_ValidDocumentNumber(t *testing.T) {
	tests := []struct {
		docNum   string
		expected inventory.DocumentNumber
	}{
		{
			docNum:   "32158465",
			expected: inventory.DocumentNumber("32158465"),
		},
		{
			docNum:   "123645659-3",
			expected: inventory.DocumentNumber("123645659-3"),
		},
	}

	for _, test := range tests {
		msg := fmt.Sprintf("valid document number %s", test.docNum)
		t.Run(msg, func(tt *testing.T) {
			result, err := inventory.NewDocumentNumber(test.docNum)
			assert.Nil(tt, err, "no error is expected")
			assert.Equal(tt, test.expected, result)
		})
	}
}

func TestDocType_ValidValues(t *testing.T) {

	tests := []struct {
		doctype  string
		expected inventory.DocumentType
	}{
		{doctype: "CC", expected: inventory.DocTypeCC},
		{doctype: "NIT", expected: inventory.DocTypeNit},
		{doctype: "CE", expected: inventory.DocTypeCE},
	}

	for _, v := range tests {
		testName := fmt.Sprintf("given the valid doctype %s no error is returned", v.doctype)
		t.Run(testName, func(tt *testing.T) {
			result, err := inventory.NewDocumentType(v.doctype)

			assert.Nil(t, err, "no error expected")
			assert.Equal(t, v.expected, result)
		})
	}

}
