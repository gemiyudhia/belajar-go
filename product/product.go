package product

import (
	"fmt"

	"github.com/google/uuid"
)

type Product struct {
	ID    uuid.UUID
	Name  string
	Price int
}

type Service struct {
	Products []Product
}

func ShowProducts(products []Product) {
	for _, product := range products {
		fmt.Println(product.Name, "-", product.Price)
	}
}

func CalculateTotal(products []Product) int {
	total := 0

	for _, product := range products {
		total += product.Price
	}

	return total
}

func FindProduct(products []Product, id uuid.UUID) (*Product, error) {
	for i := range products {
		if products[i].ID == id {
			return &products[i], nil
		}
	}

	return nil, fmt.Errorf("Product dengan ID %d tidak ditemukan.", id)
}

func (s Service) FindByID(id uuid.UUID) (*Product, error) {
	for i := range s.Products {
		if s.Products[i].ID == id {
			return &s.Products[i], nil
		}
	}

	return nil, fmt.Errorf("produk dengan ID %s tidak ditemukan", id)
}

func (s Service) GetAll() []Product {
	return s.Products
}

