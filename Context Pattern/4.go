package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	fmt.Println("Doing an HTTP request...")

	result := make(chan string)

	start := time.Now()
	go DoFirstHttpRequest(result)
	go DoFirstHttpRequest(result)
	go DoFirstHttpRequest(result)
	go DoFirstHttpRequest(result)
	go DoFirstHttpRequest(result)
	msg := <-result

	result2 := make(chan string)
	go DoSecondHttpRequest(result2, msg)
	go DoSecondHttpRequest(result2, msg)
	go DoSecondHttpRequest(result2, msg)
	go DoSecondHttpRequest(result2, msg)
	go DoSecondHttpRequest(result2, msg)
	msg2 := <-result2
	elapsed := time.Since(start)

	log.Printf("The whole HTTP requests took %s", elapsed)
	log.Printf("The whole HTTP response was: %s", msg2)

	time.Sleep(3 * time.Second)
}

// DoFirstHttpRequest performs an new HTTP request
// that can take between 0 and 500ms to be done
func DoFirstHttpRequest(result chan<- string) {
	// Do an HTTP request synchronously
	select {
	case result <- "Hellow 1":
	case <-time.After(1 * time.Second):
	}
}

// DoSecondHttpRequest performs an new HTTP request
// that can take between 0 and 500ms to be done
func DoSecondHttpRequest(result chan<- string, msg string) {
	// Do an HTTP request synchronously using msg
	select {
	case result <- "Hellow 2":
	case <-time.After(1 * time.Second):
	}
}
