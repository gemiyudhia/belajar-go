package main

import "fmt"

type Product struct {
	ID    int
	Name  string
	Price int
}

type CartItem struct {
	Product  Product
	Quantity int
}

func (item CartItem) subtotal() int {
	return item.Product.Price * item.Quantity
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

func findProduct(products []Product, id int) *Product {
	for i := range products {
		if products[i].ID == id {
			return &products[i]
		}
	}

	return nil
}

func addToCart(cart *[]CartItem, product *Product, quantity int) {
	for i := range *cart {
		if (*cart)[i].Product.ID == product.ID {
			(*cart)[i].Quantity += quantity
			return
		}
	}

	item := CartItem{
		Product:  *product,
		Quantity: quantity,
	}

	*cart = append(*cart, item)
}

func calculateCartTotal(cart []CartItem) int {
	total := 0
	for _, item := range cart {
		total += item.subtotal()
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

	var cart []CartItem

mainLoop:
	for {
		fmt.Println()
		fmt.Println("=== KASIRKU ===")
		fmt.Println("1. Lihat Produk")
		fmt.Println("2. Tambah Keranjang")
		fmt.Println("3. Total Semua Produk")
		fmt.Println("4. Keluar")
		fmt.Println("Pilih: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			showProducts(products)

		case 2:
			showProducts(products)

			fmt.Println("Masukan ID Produk: ")
			var ProductID int
			fmt.Scan(&ProductID)

			product := findProduct(products, ProductID)

			if product == nil {
				fmt.Println("Produk tidak ditemukan")
				continue
			}

			fmt.Println("Masukna Jumlah Produk: ")
			var quantity int
			fmt.Scan(&quantity)

			addToCart(&cart, product, quantity)

			fmt.Println("Produk berhasil ditambahkan ke keranjang")

		case 3:
			fmt.Println("=== KERANJANG ===")

			if len(cart) == 0 {
				fmt.Println("Keranjang masih kosong")
				continue
			}

			for _, item := range cart {
				fmt.Println(item.Product.Name, "x", item.Quantity, "=", item.subtotal())
			}

		case 4:
			fmt.Println("Terima Kasih")
			break mainLoop

		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}
