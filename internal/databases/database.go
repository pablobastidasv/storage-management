package databases

import "database/sql"

func New() (*sql.DB, error) {
	return newSqliteDatabase()
}
