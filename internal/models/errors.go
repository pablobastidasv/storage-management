package models

type DomainError struct {
	desc string
}

func (e *DomainError) Error() string {
	return e.desc
}

func NewDomainError(description string) *DomainError {
	return &DomainError{
		desc: description,
	}
}
