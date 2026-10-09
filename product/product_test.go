package product

import (
	"testing"

	"github.com/google/uuid"
)

type FakeRepository struct {
	Products []Product
}

func (r *FakeRepository) GetAll() ([]Product, error) {
	return r.Products, nil
}

func (r *FakeRepository) FindByID(id uuid.UUID) (*Product, error) {
	for i := range r.Products {
		if r.Products[i].ID == id {
			return &r.Products[i], nil
		}
	}

	return nil, nil
}

func (r *FakeRepository) Create(product Product) (Product, error) {
	r.Products = append(r.Products, product)
	return product, nil
}

func (r *FakeRepository) Update(product Product) error {
	for i := range r.Products {
		if r.Products[i].ID == product.ID {
			r.Products[i] = product
			return nil
		}
	}

	return nil
}

func (r *FakeRepository) Delete(id uuid.UUID) error {
	for i := range r.Products {
		if r.Products[i].ID == id {
			r.Products = append(
				r.Products[:i],
				r.Products[i+1:]...,
			)
			return nil
		}
	}

	return nil
}

func TestServiceCreateWithFakeRepository(t *testing.T) {
	repository := &FakeRepository{}

	service := Service{
		Repository: repository,
	}

	product, err := service.Create("Indomie", 3500)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if product.Name != "Indomie" {
		t.Errorf("expected Indomie, got %s", product.Name)
	}

	if len(repository.Products) != 1 {
		t.Errorf(
			"expected 1 product, got %d",
			len(repository.Products),
		)
	}
}
