package postgres

const (
	sqlClientTable = "clients"
)

type sqlClient struct {
	DocType   string `db:"doc_type"`
	DocNumber string `db:"doc_number"`
	Name      string `db:"name"`
}
