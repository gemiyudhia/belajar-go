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

func producer(ch chan<- int) {
	ch <- 100
	ch <- 200
	ch <- 300

	close(ch)
}

func consumer(ch <-chan int) {
	for total := range ch {
		fmt.Println("Total: ", total)
	}
}

func main() {
	ch := make(chan int)

	go producer(ch)

	consumer(ch)

	fmt.Println("Selesai")
}
