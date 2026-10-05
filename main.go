package main

import (
	"fmt"
	"kasirku-go/product"

	"github.com/google/uuid"
)

func main() {
	service := product.Service{
		Products: []product.Product{
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
		},
	}

	fmt.Println("=== AWAL ===")

	for _, p := range service.GetAll() {
		fmt.Println(p.Name, p.Price)
	}

	fmt.Println("=== CREATE ===")

	newProduct := service.Create("Kopi Hitam", 2000)

	fmt.Println(newProduct.Name, newProduct.Price)

	fmt.Println("=== UPDATE ===")
	err := service.Update(
		newProduct.ID,
		"Kopi Susu",
		5000,
	)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	found, err := service.FindByID(newProduct.ID)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println(found.Name, found.Price)

	fmt.Println("=== DELETE ===")

	err = service.Delete(newProduct.ID)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println("Produk berhasil disimpan")
	fmt.Println("=== AKHIR ===")

	for _, p := range service.GetAll() {
		fmt.Println(p.Name, p.Price)
	}

}
