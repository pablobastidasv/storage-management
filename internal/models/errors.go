package models

import (
	"strings"
)

type (
	DomainError struct {
		desc string
	}

	ValidationError struct {
		Messages []string
	}
)

func (e *ValidationError) Error() string {
	return strings.Join(e.Messages, ";")
}

func (e *DomainError) Error() string {
	return e.desc
}

func NewDomainError(description string) *DomainError {
	return &DomainError{
		desc: description,
	}
}

func NewValidationError(messages []string) *ValidationError {
    return &ValidationError{
        Messages: messages,
    }
}
