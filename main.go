package main

import "fmt"

type Product struct {
	ID    int
	Name  string
	Price int
}

func showProducts(products []Product) {
	for _, product := range products {
		fmt.Println(product.Name, "-", product.Price)
	}
}

func calculateTotal(products []Product) int {
	total := 0

	for _, product := range products {
		total += product.Price
	}

	return total
}

func main() {
	products := []Product{
		{
			ID:    1,
			Name:  "Indomie Goreng",
			Price: 2000,
		},
		{
			ID:    2,
			Name:  "Teh Botol",
			Price: 1000,
		},
		{
			ID:    3,
			Name:  "Kopi",
			Price: 3000,
		},
	}
	showProducts(products)

	total := calculateTotal(products)
	fmt.Println("Total: ", total)
}
