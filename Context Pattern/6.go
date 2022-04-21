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
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)

	start := time.Now()
	go DoHttpRequest4(ctx, result)
	go DoHttpRequest4(ctx, result)
	go DoHttpRequest4(ctx, result)
	go DoHttpRequest4(ctx, result)
	go DoHttpRequest4(ctx, result)
	msg := <-result
	cancel()

	elapsed := time.Since(start)

	log.Printf("The whole HTTP requests took %s", elapsed)
	log.Printf("The whole HTTP response was: %s", msg)

	time.Sleep(3 * time.Second)
}

// DoHttpRequest performs an new HTTP request
// that can take between 0 and 500ms to be done
func DoHttpRequest4(ctx context.Context, result chan<- string) {
	// Do an HTTP request synchronously
	select {
	case result <- "Hello World":
	case <-ctx.Done():
	}
}
