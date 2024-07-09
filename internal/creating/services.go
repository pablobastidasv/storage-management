package creating

import "context"

//go:generate mockery --case=snake --outpkg=creatingmocks --output=./creatingmocks --name ClientCreator
type ClientCreator interface {
    Create(ctx context.Context, docType, docNumber, name string) error
}

