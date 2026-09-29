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

mainLoop:
	for {
		fmt.Println()
		fmt.Println("=== KASIRKU ===")
		fmt.Println("1. Lihat Produk")
		fmt.Println("2. Total Semua Produk")
		fmt.Println("3. Keluar")
		fmt.Println("Pilih: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			showProducts(products)

		case 2:
			total := calculateTotal(products)
			fmt.Println("Total: ", total)

		case 3:
			fmt.Println("Terima Kasih!")
			break mainLoop

		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}
