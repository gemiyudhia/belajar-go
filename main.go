package main

import (
	"fmt"
	"kasirku-go/product"

	"github.com/google/uuid"
)

func main() {
	products := []product.Product{
		{
			ID:    uuid.New(),
			Name:  "Indomie Goreng",
			Price: 2000,
		},
		{
			ID:    uuid.New(),
			Name:  "Teh Botol",
			Price: 1000,
		},
		{
			ID:    uuid.New(),
			Name:  "Kopi",
			Price: 3000,
		},
	}
	service := product.Service{
		Products: products,
	}

	for _, p := range products {
		fmt.Println(p.ID, "-", p.Name, "-", p.Price)
	}

	productID := products[0].ID

	found, err := service.FindByID(productID)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	products = service.GetAll()

	for _, p := range products {
		fmt.Println(p.Name, p.Price)
	}

	fmt.Println("Produk ditemukan:")
	fmt.Println(found.Name)
	fmt.Println(found.Price)

}
