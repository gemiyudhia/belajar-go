package product

import "testing"

func TestCreateProduct(t *testing.T) {
	service := Service{}

	product, err := service.Create("Indomie", 3000)

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if product.Name != "Indomie" {
		t.Errorf("Expected name Indomie, got %s", product.Name)
	}

	if product.Price != 3000 {
		t.Errorf("Expected price 3000, got %d", product.Price)
	}

	if len(service.Products) != 1 {
		t.Errorf("Expected 1 product, got %d", len(service.Products))
	}
}

func TestCreateProductValidation(t *testing.T) {
	service := Service{}

	_, err := service.Create("", 3500)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if err != ErrEmptyName {
		t.Errorf("expected ErrEmptyName, got %v", err)
	}
}

func TestCreateProductInvalidPrice(t *testing.T) {
	service := Service{}

	_, err := service.Create("Indomie", 0)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if err != ErrInvalidPrice {
		t.Errorf("expected ErrInvalidPrice, got %v", err)
	}
}

func TestFindProduct(t *testing.T) {
	service := Service{}

	product, err := service.Create("Indomie", 3500)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	found, err := service.FindByID(product.ID)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if found.Name != "Indomie" {
		t.Errorf("expected Indomie, got %s", found.Name)
	}
}

func TestDeleteProduct(t *testing.T) {
	service := Service{}

	product, err := service.Create("Indomie", 3500)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	err = service.Delete(product.ID)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(service.Products) != 0 {
		t.Errorf("expected 0 products, got %d", len(service.Products))
	}
}