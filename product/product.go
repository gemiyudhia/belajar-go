package product

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

var (
	ErrEmptyName    = errors.New("nama produk tidak boleh kosong")
	ErrInvalidPrice = errors.New("harga produk harus lebih dari 0")
)

type Product struct {
	ID    uuid.UUID
	Name  string
	Price int
}

type Service struct {
	Repository Repository
}

func validateProduct(name string, price int) error {
	if name == "" {
		return ErrEmptyName
	}

	if price <= 0 {
		return ErrInvalidPrice
	}

	return nil
}
func FindProduct(products []Product, id uuid.UUID) (*Product, error) {
	for i := range products {
		if products[i].ID == id {
			return &products[i], nil
		}
	}

	return nil, fmt.Errorf("Product dengan ID %d tidak ditemukan.", id)
}

func (s *Service) FindByID(id uuid.UUID) (*Product, error) {
	return s.Repository.FindByID(id)
}

func (s *Service) GetAll() []Product {
	return s.Repository.GetAll()
}

func (s *Service) Create(name string, price int) (Product, error) {
	err := validateProduct(name, price)

	if err != nil {
		return Product{}, err
	}

	newProduct := Product{
		ID:    uuid.New(),
		Name:  name,
		Price: price,
	}

	product := s.Repository.Create(newProduct)

	return product, nil
}

func (s *Service) Update(id uuid.UUID, name string, price int) error {
	err := validateProduct(name, price)

	if err != nil {
		return err
	}

	product := Product{
		ID:    id,
		Name:  name,
		Price: price,
	}

	return s.Repository.Update(product)
}

func (s *Service) Delete(id uuid.UUID) error {
	return s.Repository.Delete(id)

}
