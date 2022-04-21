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
	go DoFirstHttpRequest1(ctx, result)
	go DoFirstHttpRequest1(ctx, result)
	go DoFirstHttpRequest1(ctx, result)
	go DoFirstHttpRequest1(ctx, result)
	go DoFirstHttpRequest1(ctx, result)
	msg := <-result

	result2 := make(chan string)
	ctx = context.WithValue(ctx, "msg", msg)
	go DoSecondHttpRequest1(ctx, result2)
	go DoSecondHttpRequest1(ctx, result2)
	go DoSecondHttpRequest1(ctx, result2)
	go DoSecondHttpRequest1(ctx, result2)
	go DoSecondHttpRequest1(ctx, result2)
	msg2 := <-result2
	cancel()
	elapsed := time.Since(start)

	log.Printf("The whole HTTP requests took %s", elapsed)
	log.Printf("The whole HTTP response was: %s", msg2)

	time.Sleep(3 * time.Second)
}

func DoFirstHttpRequest1(ctx context.Context, result chan<- string) {
	// Do an HTTP request synchronously
	select {
	case result <- "Hellow 1":
	case <-ctx.Done():
	}
}

func DoSecondHttpRequest1(ctx context.Context, result chan<- string) {
	// Do an HTTP request synchronously using the ctx variables
	select {
	case result <- "Hellow 2":
	case <-ctx.Done():
	}
}
