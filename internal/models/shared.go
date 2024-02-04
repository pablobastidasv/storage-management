package models

import (
	"strings"
)

type ProductId string

func ProductIdFrom(value string) (ProductId, error) {
	if len(strings.Trim(value, " ")) == 0 {
		return ProductId("INVALID"), NewDomainError("product id cannot be empty")
	}
	trimmed := strings.Trim(value, " ")
	return ProductId(trimmed), nil
}

func (p *ProductId) ToString() string {
	return string(*p)
}
