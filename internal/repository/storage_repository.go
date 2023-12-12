package repository

import (
	"co.bastriguez/inventory/internal/models"
	"context"
	"database/sql"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type StorageRepository interface {
	FetchItemsByStorage(ctx context.Context, storage *models.Storage) ([]models.InventoryItem, error)
	FindItemByProductId(ctx context.Context, storageId string, productId string) (*models.InventoryItem, error)
	UpsertItem(ctx context.Context, storageId string, item *models.InventoryItem) error
	FindMainStorage(ctx context.Context) (*models.Storage, error)
}

// === mongo implementation

const InventoryItemsCollectionName = "inventoryItems"

type mongoRepository struct {
	collection *mongo.Collection
}

func (m mongoRepository) FetchItemsByStorage(ctx context.Context, _ *models.Storage) ([]models.InventoryItem, error) {
	found, err := m.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	var items []models.InventoryItem
	if err := found.All(ctx, &items); err != nil {
		return nil, err
	}

	return items, nil
}

func (m mongoRepository) FindItemByProductId(ctx context.Context, _ string, productId string) (*models.InventoryItem, error) {
	res := m.collection.FindOne(ctx, bson.D{{"product.id", productId}})
	var repoItem InventoryItem
	if err := res.Decode(&repoItem); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}

	return &models.InventoryItem{
		Product: models.InventoryProduct{
			Id:           repoItem.Product.Id,
			Name:         repoItem.Product.Name,
			Presentation: repoItem.Product.Presentation,
		},
		Qty: repoItem.Qty,
	}, nil
}

// UpsertItem TODO: if item does not, exist create a new one
func (m mongoRepository) UpsertItem(ctx context.Context, _ string, item *models.InventoryItem) error {
	filter := bson.D{{"product.id", item.Product.Id}}

	prod := InventoryProduct{
		Id:           item.Product.Id,
		Name:         item.Product.Name,
		Presentation: item.Product.Presentation,
	}
	update := bson.D{{"$set", bson.D{{"qty", item.Qty}, {"product", prod}}}}
	opts := options.Update().SetUpsert(true)

	_, err := m.collection.UpdateOne(ctx, filter, update, opts)

	return err
}

func (m mongoRepository) FindMainStorage(_ context.Context) (*models.Storage, error) {
	return &models.Storage{
		Id: "313fbcf2-daeb-405d-b9e6-94649a33c5f2",
	}, nil
}

// sql implementatioon
type relationalRepository struct {
	db *sql.DB
}

func (r *relationalRepository) FindMainStorage(_ context.Context) (*models.Storage, error) {
	row := r.db.QueryRow("select s.id from storages s limit 1")
	var storage models.Storage
	err := row.Scan(&storage.Id)
	if err != nil {
		return nil, err
	}

	return &storage, nil
}

func (r *relationalRepository) FindItemByProductId(_ context.Context, storageId string, productId string) (*models.InventoryItem, error) {
	row := r.db.QueryRow(
		`select i.quantity, p.name, p.presentation from items i 
			join public.products p on p.id = i.product_id 
			where i.storage_id = $1 
			and i.product_id = $2`,
		storageId,
		productId,
	)

	var item models.InventoryItem
	var presentation string
	err := row.Scan(&item.Qty, &item.Product.Name, &presentation)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	item.Product.Id = productId
	item.Product.Presentation = models.NewPresentation(presentation)

	return &item, nil
}

func (r *relationalRepository) UpsertItem(_ context.Context, storageId string, item *models.InventoryItem) error {
	_, err := r.db.Exec(`
		INSERT INTO items(storage_id, product_id, quantity)
			VALUES($1, $2, $3) 
			ON CONFLICT (storage_id, product_id) 
			DO 
			   UPDATE SET quantity = $3;
		`, storageId, item.Product.Id, item.Qty)
	if err != nil {
		return err
	}

	return nil
}

func (r *relationalRepository) FetchItemsByStorage(_ context.Context, _ *models.Storage) ([]models.InventoryItem, error) {
	rows, err := r.db.Query(`select i.quantity, p.id, p.name, p.presentation 
								 from items i 
								 join public.products p on p.id = i.product_id
								 order by p.name`)
	if err != nil {
		return nil, nil
	}
	defer rows.Close()

	var items []models.InventoryItem
	for rows.Next() {
		var item models.InventoryItem
		var presentation string
		if err := rows.Scan(&item.Qty, &item.Product.Id, &item.Product.Name, &presentation); err != nil {
			return nil, nil
		}
		item.Product.Presentation = models.NewPresentation(presentation)

		items = append(items, item)
	}

	return items, nil
}

func NewStorageMongoRepository(database *mongo.Database) StorageRepository {
	collection := database.Collection(InventoryItemsCollectionName)
	return &mongoRepository{
		collection: collection,
	}
}

func NewStorageRelationalRepository(db *sql.DB) StorageRepository {
	return &relationalRepository{
		db: db,
	}
}
