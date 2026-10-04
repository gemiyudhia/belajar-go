package main

import (
	"context"
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

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker dihentikan")
			return

		default:
			fmt.Println("Worker sedang bekerja")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go worker(ctx)

	time.Sleep(2 * time.Second)

	cancel()

	time.Sleep(1 * time.Second)
}
