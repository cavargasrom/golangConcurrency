package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	fmt.Println("Doing an HTTP request...")

	result := make(chan string)

	ctx, cancel := context.WithCancel(context.Background())

	start := time.Now()
	go DoHttpRequest3(ctx, result, 1)
	go DoHttpRequest3(ctx, result, 2)
	go DoHttpRequest3(ctx, result, 3)
	go DoHttpRequest3(ctx, result, 4)
	go DoHttpRequest3(ctx, result, 5)
	msg := <-result
	cancel()

	elapsed := time.Since(start)

	log.Printf("The whole HTTP requests took %s", elapsed)
	log.Printf("The whole HTTP response was: %s", msg)

	time.Sleep(3 * time.Second)
}

// DoHttpRequest performs an new HTTP request
// that can take between 0 and 500ms to be done
func DoHttpRequest3(ctx context.Context, result chan<- string, i int) {
	// Do an HTTP request synchronously
	select {
	case result <- "Hello World":
		fmt.Printf("Goroutine finished #%d\n", i)
	case <-ctx.Done():
		fmt.Printf("Goroutine finished #%d\n", i)
	}
}
