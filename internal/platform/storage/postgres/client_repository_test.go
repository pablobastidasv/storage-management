package postgres_test

import (
	"context"
	"errors"
	"testing"

	"co.bastriguez/inventory/internal/inventory"
	"co.bastriguez/inventory/internal/platform/storage/postgres"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tinygg/gofaker"
)

func Test_ClientRepository_Save(t *testing.T) {
	t.Run("Repository Error", func(t *testing.T) {
		docType, docNum, name := "CC", "1234567890", gofaker.Name()
		address, email, phone := gofaker.Address().Address, gofaker.Email(), gofaker.Phone()
		client, err := inventory.NewClient(docType, docNum, name, address, email, phone)
		require.NoError(t, err, "error when creating client")

		db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		require.NoError(t, err, "error creating sqlmock")

		sqlMock.ExpectExec(
			"INSERT INTO clients (doc_type, doc_number, name, address, email, phone) VALUES (?, ?, ?, ?, ?, ?)").
			WithArgs(docType, docNum, name, address, email, phone).
			WillReturnError(errors.New("something-failed"))

		repo := postgres.NewClientRepository(db)

		err = repo.Save(context.Background(), client)

		assert.NoError(t, sqlMock.ExpectationsWereMet())
		assert.Error(t, err)
	})

	t.Run("Succeed", func(t *testing.T) {


		docType, docNum, name := "CC", "1234567890", "Pepito Carabali"
		address, email, phone := gofaker.Address().Address, gofaker.Email(), gofaker.Phone()
		client, err := inventory.NewClient(docType, docNum, name, address, email, phone)
		require.NoError(t, err, "error when creating client")

		db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		require.NoError(t, err, "error creating sqlmock")

		sqlMock.ExpectExec(
			"INSERT INTO clients (doc_type, doc_number, name, address, email, phone) VALUES (?, ?, ?, ?, ?, ?)").
			WithArgs(docType, docNum, name, address, email, phone).
			WillReturnResult(sqlmock.NewResult(0, 1))

		repo := postgres.NewClientRepository(db)


		err = repo.Save(context.Background(), client)

		assert.NoError(t, sqlMock.ExpectationsWereMet())
		assert.NoError(t, err)

	})

}
