package product

import "fmt"

type Product struct {
	ID    int
	Name  string
	Price int
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

func FindProduct(products []Product, id int) (*Product, error) {
	for i := range products {
		if products[i].ID == id {
			return &products[i], nil
		}
	}

	return nil, fmt.Errorf("Product dengan ID %d tidak ditemukan.", id)
}