package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"co.bastriguez/inventory/internal/inventory"
	"github.com/huandu/go-sqlbuilder"
)

type ClientRepository struct {
	db *sql.DB
	// dbTimout time.Duration
}

func NewClientRepository(db *sql.DB) *ClientRepository {
	return &ClientRepository{
		db: db,
	}
}

// Save implements the inventory.Persist interface
func (r *ClientRepository) Save(c context.Context, client inventory.Client) error {
	clientSql := sqlbuilder.NewStruct(new(sqlClient))
	query, args := clientSql.InsertInto(sqlClientTable, sqlClient{
		DocType:   string(client.Document.DocumentType),
		DocNumber: string(client.Document.Number),
		Name:      string(client.Name),
		Address:   string(client.Address),
		Email:     string(client.Email),
		Phone:     string(client.PhoneNumber),
	}).Build()

	// ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	// defer cancel()

	// _, err := r.db.ExecContext(ctxTimeout, query, args...)
	_, err := r.db.ExecContext(c, query, args...)
	if err != nil {
		return fmt.Errorf("error trying to persist the client on database: %v", err)
	}

	return nil
}
