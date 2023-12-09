package repository

import "co.bastriguez/inventory/internal/models"

type product struct {
	Id           string              `bson:"_id"`
	Name         string              `bson:"name"`
	Presentation models.Presentation `bson:"presentation"`
}
