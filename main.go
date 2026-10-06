package main

import (
	"fmt"
	"kasirku-go/product"
)

func main() {
	repository := &product.MemoryRepository{}

	service := product.Service{
		Repository: repository,
	}

	fmt.Println("=== CREATE ===")

	newProduct, err := service.Create("Indomie", 3500)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println(newProduct.Name, newProduct.Price)

	fmt.Println("\n === GET ALL ===")

	product := service.GetAll()

	for _, p := range product {
		fmt.Println(p.Name, p.Price)
	}

}
