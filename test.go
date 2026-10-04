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
	paymentCh := make(chan string)
	reportCh := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		paymentCh <- "Pembayaran Berhasil"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		reportCh <- "Laporan Selesai"
	}()

	for i := 0; i < 2; i++ {
		select {
		case message := <-paymentCh:
			fmt.Println(message)

		case message := <-reportCh:
			fmt.Println(message)
		}
	}

}
