package storage

import "database/sql"

type Storage interface {
}

type storage struct {
	db *sql.DB
}

func New(db *sql.DB) (Storage, error) {
	return &storage{
		db: db,
	}, nil
}
