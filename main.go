package main

import (
	"fmt"
	"kasirku-go/cart"
	"kasirku-go/payment"
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

	items := []cart.Item{
		{
			Product:  products[0],
			Quantity: 2,
		},
		{
			Product:  products[1],
			Quantity: 1,
		},
	}

	fmt.Println("=== KASIRKU ===")

	for _, item := range items {
		fmt.Println(
			item.Product.Name,
			"X",
			item.Quantity,
			"=",
			item.Subtotal(),
		)
	}

	total := cart.Total(items)

	fmt.Println("Total: ", total)

	debit := payment.DebitPayment{
		CardNumber: 123,
	}

	service := payment.PaymentService{
		PaymentMethod: debit,
	}

	err := service.Pay(total)

	if err != nil {
		fmt.Println("Error", err)
	}

}
