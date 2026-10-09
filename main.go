package main

import (
	"fmt"
	"kasirku-go/database"
	"kasirku-go/product"

	"github.com/google/uuid"
)

func main() {
	db, err := database.Connect()

	if err != nil {
		fmt.Println("Database error: ", err)
		return
	}

	defer db.Close()

	repository := &product.MySQLRepository{
		DB: db,
	}

	service := product.Service{
		Repository: repository,
	}

	newProduct, err := service.Create("Indomie", 3500)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println("ID Baru: ", newProduct.ID)

	found, err := service.FindByID(newProduct.ID)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	randomID := uuid.New()

	_, err = service.FindByID(randomID)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("\n === UPDATE ===")

	newProduct.Name = "Indomie Goreng"
	newProduct.Price = 5000

	err = service.Update(
		newProduct.ID,
		newProduct.Name,
		newProduct.Price,
	)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println("Produk berhasil di update")

	updateProduct, err := service.FindByID(newProduct.ID)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println("\n === DELETE ===")

	err = service.Delete(newProduct.ID)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println("produk berhasil dihapus")

	fmt.Println("ID: ", updateProduct.ID)
	fmt.Println("Nama: ", updateProduct.Name)
	fmt.Println("Harga: ", updateProduct.Price)

	fmt.Println("Produk Ditemukan")
	fmt.Println("ID", found.ID)
	fmt.Println("Nama: ", found.Name)
	fmt.Println("Harga: ", found.Price)

	fmt.Println("\n=== SEMUA PRODUK ===")

	products, err := service.GetAll()

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	for _, p := range products {
		fmt.Println(
			p.ID,
			"|",
			p.Name,
			"|",
			p.Price,
		)
	}
}
