package main

import (
	"errors"
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

	fmt.Println("\n=== UJI COBA VALIDASI ERROR ===")
	// Menguji input validasi yang gagal tanpa mengganggu alur utama
	_, errInvalid := service.Create("", 0)
	if errors.Is(errInvalid, product.ErrEmptyName) {
		fmt.Println("Validasi berhasil: Nama produk wajib diisi")
	} else if errInvalid != nil {
		fmt.Println("Error validasi:", errInvalid)
	}

	fmt.Println("\n=== CREATE ===")
	newProduct, err := service.Create("", -2000)
	if err != nil {
		fmt.Println("Error saat membuat produk:", err)
		return
	}
	fmt.Println("Berhasil dibuat:", newProduct.Name, newProduct.Price)

	fmt.Println("\n=== UPDATE ===")
	err = service.Update(
		newProduct.ID,
		"Kopi Susu",
		0,
	)
	if err != nil {
		fmt.Println("Error saat update:", err)
		return
	}

	found, err := service.FindByID(newProduct.ID)
	if err != nil {
		fmt.Println("Error saat mencari produk:", err)
		return
	}
	fmt.Println("Setelah diupdate:", found.Name, found.Price)

	fmt.Println("\n=== DELETE ===")
	err = service.Delete(newProduct.ID)
	if err != nil {
		fmt.Println("Error saat menghapus produk:", err)
		return
	}
	fmt.Println("Produk berhasil dihapus")

	fmt.Println("\n=== AKHIR ===")
	for _, p := range service.GetAll() {
		fmt.Println(p.Name, p.Price)
	}
}