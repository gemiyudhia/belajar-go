package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 1; i <= 3; i++ {
		fmt.Println(name, i)
		time.Sleep(500 * time.Millisecond)
	}
}

func producer(ch chan string) {
	ch <- "Indomie"
	ch <- "Teh Botol"
	ch <- "Kopi"

	close(ch)
}

func main() {
	ch := make(chan string)

	go producer(ch)

	for product := range ch {
		fmt.Println("Produk: ", product)
	}

	fmt.Println("Selesai")
}
