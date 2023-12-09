package repository

import (
	"co.bastriguez/inventory/internal/databases"
	"co.bastriguez/inventory/internal/models"
	"context"
	"database/sql"
	"log"
	"reflect"
	"testing"
)

// Used to debug database connection, there is not a mocked database therefore this test is not reliable
func _Test_repository_RetrieveItemByStorage(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		storage *models.Storage
		ctx     context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []models.InventoryItem
		wantErr bool
	}{
		{
			name: "it should work",
			fields: fields{
				db: testDatabase(),
			},
			args: args{
				storage: nil,
				ctx:     context.Background(),
			},
			want:    []models.InventoryItem{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{
				db: tt.fields.db,
			}
			got, err := r.FetchItemsByStorage(tt.args.ctx, tt.args.storage)
			if (err != nil) != tt.wantErr {
				t.Errorf("RetrieveItemByStorage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RetrieveItemByStorage() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func testDatabase() *sql.DB {
	db, err := databases.New()
	if err != nil {
		log.Fatalf("there was an error connecting to the database, %s", err.Error())
	}
	return db
}
