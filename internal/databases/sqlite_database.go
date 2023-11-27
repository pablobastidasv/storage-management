package databases

import (
	"database/sql"
	"os"
)

func newSqliteDatabase() (*sql.DB, error) {
	path := "./bastriguez.db"
	if _, err := os.Stat(path); err != nil {
		file, err := os.Create(path)
		if err != nil {
			return nil, err
		}
		err = file.Close()
		if err != nil {
			return nil, err
		}
	}

	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	return db, nil
}
