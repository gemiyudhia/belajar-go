package product

import (
	"fmt"

	"github.com/google/uuid"
)

type Repository interface {
	GetAll() ([]Product, error)
	FindByID(id uuid.UUID) (*Product, error)
	Create(product Product) (Product, error)
	Update(product Product) error
	Delete(id uuid.UUID) error
}

type MemoryRepository struct {
	Products []Product
}

func (r *MemoryRepository) GetAll() ([]Product, error) {
	return r.Products, nil
}

func (r *MemoryRepository) FindByID(id uuid.UUID) (*Product, error) {
	for i := range r.Products {
		if r.Products[i].ID == id {
			return &r.Products[i], nil
		}
	}

	return nil, fmt.Errorf(
		"produk dengan id %s tidak ditemukan",
		id,
	)
}

func (r *MemoryRepository) Create(product Product) (Product, error) {
	r.Products = append(r.Products, product)

	return product, nil
}

func (r *MemoryRepository) Update(product Product) error {
	for i := range r.Products {
		if r.Products[i].ID == product.ID {
			r.Products[i] = product
			return nil
		}
	}

	return fmt.Errorf(
		"produk dengan ID %s tidak ditemukan",
		product.ID,
	)
}

func (r *MemoryRepository) Delete(id uuid.UUID) error {
	for i := range r.Products {
		if r.Products[i].ID == id {
			r.Products = append(r.Products[:i], r.Products[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf(
		"Produk dengan ID %s tidak ditemukan",
		id,
	)
}
