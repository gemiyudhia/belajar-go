package main

import (
	"errors"
	"fmt"
)

type ProductNotFoundError struct {
	ProductID int
}

type error interface {
	Error() string
}

func (e ProductNotFoundError) Error() string {
	return fmt.Sprintf(
		"produk dengan id %d tidak ditemukan",
		e.ProductID,
	)
}

func findProduct(id int) error {
	if id != 1 {
		return ProductNotFoundError{
			ProductID: id,
		}
	}

	return nil
}

func main() {
	err := findProduct(10)

	if err != nil {
		fmt.Println("Error: ", err)

		var productNotFoundError ProductNotFoundError

		if errors.As(err, &productNotFoundError) {
			fmt.Println("Jenis error: ProductNotFoundError")
			fmt.Println("ID produk: ", productNotFoundError.ProductID)
		}

		return
	}

	fmt.Println("Produk Ditemukan")
}
