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

func findProduct(products []Product, id int) (*Product, error) {
	for i := range products {
		if products[i].ID == id {
			return &products[i], nil
		}
	}

	return nil, fmt.Errorf("Product dengan ID %d tidak ditemukan.", id)
}

func addToCart(cart *[]CartItem, product *Product, quantity int) error {
	if quantity <= 0 {
		return fmt.Errorf("Jumlah produk harus lebih dari 0")
	}

	for i := range *cart {
		if (*cart)[i].Product.ID == product.ID {
			(*cart)[i].Quantity += quantity
			return nil
		}
	}

	item := CartItem{
		Product:  *product,
		Quantity: quantity,
	}

	*cart = append(*cart, item)

	return nil
}

func calculateCartTotal(cart []CartItem) int {
	total := 0
	for _, item := range cart {
		total += item.subtotal()
	}
	return total
}

type PaymentMethod interface {
	Pay(total int) error
}

type CashPayment struct {
	Cash int
}

func (c CashPayment) Pay(total int) error {
	if c.Cash < total {
		return fmt.Errorf("Uang anda tidak cukup.")
	}

	change := c.Cash - total

	fmt.Println("Pembayaran Berhasil!")
	fmt.Println("Kembalian: ", change)

	return nil
}

type QrisPayment struct{}

func (q QrisPayment) Pay(total int) error {
	fmt.Println("Pembayaran QRIS berhasil!")
	fmt.Println("Total: ", total)
	return nil
}

func processPayment(payment PaymentMethod, total int) error {
	return payment.Pay(total)
}

type PaymentService struct {
	paymentMethod PaymentMethod
}

func (s *PaymentService) Pay(total int) error {
	return s.paymentMethod.Pay(total)
}

type DebitPayment struct {
	CardNumber int
}

func (d DebitPayment) Pay(total int) error {
	fmt.Println("Pembayaran Debit Berhasil!")
	fmt.Println("Total: ", total)

	return nil
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

	// cash := CashPayment{
	// 	Cash: 10000,
	// }
	// qris := QrisPayment{}
	debit := DebitPayment{
		CardNumber: 1234,
	}

	service := PaymentService{
		paymentMethod: debit,
	}

	err := service.Pay(7500)

	if err != nil {
		fmt.Println("Error: ", err)
	}

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

			product, err := findProduct(products, ProductID)

			if err != nil {
				fmt.Println("Error: ", err)
				continue
			}

			fmt.Println("Masukna Jumlah Produk: ")
			var quantity int
			fmt.Scan(&quantity)

			err = addToCart(&cart, product, quantity)

			if err != nil {
				fmt.Println("Error: ", err)
				continue
			}

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
