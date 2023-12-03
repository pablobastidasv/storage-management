package repository

import (
	"co.bastriguez/inventory/internal/models"
	"database/sql"
	"reflect"
	"testing"
)

// Used to debug database connection, there is not a mocked database therefore this test is not reliable
func _Test_productRepo_FetchProducts(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	tests := []struct {
		name    string
		fields  fields
		want    []models.Product
		wantErr bool
	}{
		{
			name: "it should work",
			fields: fields{
				db: testDatabase(),
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := productRepo{
				db: tt.fields.db,
			}
			got, err := p.FetchProducts()
			if (err != nil) != tt.wantErr {
				t.Errorf("FetchProducts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FetchProducts() got = %v, want %v", got, tt.want)
			}
		})
	}
}
