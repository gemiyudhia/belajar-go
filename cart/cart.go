package cart

import (
	"fmt"
	"kasirku-go/product"
)

type Item struct {
	Product  product.Product
	Quantity int
}

func (item Item) Subtotal() int {
	return item.Product.Price * item.Quantity
}

func Total(cart []Item) int {
	total := 0

	for _, item := range cart {
		total += item.Subtotal()
	}

	return total
}

func AddToCart(cart *[]Item, product *product.Product, quantity int) error {
	if quantity <= 0 {
		return fmt.Errorf("Jumlah produk harus lebih dari 0")
	}

	for i := range *cart {
		if (*cart)[i].Product.ID == product.ID {
			(*cart)[i].Quantity += quantity
			return nil
		}
	}

	item := Item{
		Product:  *product,
		Quantity: quantity,
	}

	*cart = append(*cart, item)

	return nil
}
