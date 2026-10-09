package product

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
)

func TestMySQLRepositoryCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("gagal membuat mock database: %v", err)
	}
	defer db.Close()

	repository := &MySQLRepository{
		DB: db,
	}

	newProduct := Product{
		ID:    uuid.New(),
		Name:  "Indomie",
		Price: 3500,
	}

	mock.ExpectExec("INSERT INTO products").
		WithArgs(
			newProduct.ID.String(),
			newProduct.Name,
			newProduct.Price,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))

	createdProduct, err := repository.Create(newProduct)

	if err != nil {
		t.Fatalf("gagal membuat produk: %v", err)
	}

	if createdProduct.ID != newProduct.ID {
		t.Errorf("ID tidak sesuai")
	}

	if createdProduct.Name != newProduct.Name {
		t.Errorf(
			"nama seharusnya %s, dapat %s",
			newProduct.Name,
			createdProduct.Name,
		)
	}

	if createdProduct.Price != newProduct.Price {
		t.Errorf(
			"harga seharusnya %d, dapat %d",
			newProduct.Price,
			createdProduct.Price,
		)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("ekspektasi database tidak terpenuhi: %v", err)
	}
}

func TestMySQLRepositoryCreateError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("gagal membuat mock database: %v", err)
	}
	defer db.Close()

	repository := &MySQLRepository{
		DB: db,
	}

	newProduct := Product{
		ID:    uuid.New(),
		Name:  "Indomie",
		Price: 3500,
	}

	mock.ExpectExec("INSERT INTO products").
		WithArgs(
			newProduct.ID.String(),
			newProduct.Name,
			newProduct.Price,
		).
		WillReturnError(fmt.Errorf("database error"))

	_, err = repository.Create(newProduct)

	if err == nil {
		t.Fatal("seharusnya mengembalikan error")
	}

	if err.Error() != "database error" {
		t.Errorf("error tidak sesuai: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("ekspektasi database tidak terpenuhi: %v", err)
	}
}