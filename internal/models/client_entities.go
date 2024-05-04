package models

const (
	NIT DocumentType = "NIT"
	CC = "CC"
	CE = "CE"
)

type (
	DocumentType string

	Client struct {
		Identity Identity `bson:"_id"`
		Name     string   `bson:"name"`
	}

	Identity struct {
		DocumentType DocumentType `bson:"type"`
		Number       string       `bson:"number"`
	}
)
