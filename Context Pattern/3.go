package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Doing an HTTP request...")

	result := make(chan string)

	start := time.Now()
	go DoHttpRequest2(result, 1)
	go DoHttpRequest2(result, 2)
	go DoHttpRequest2(result, 3)
	go DoHttpRequest2(result, 4)
	go DoHttpRequest2(result, 5)
	response := <-result
	elapsed := time.Since(start)

	log.Printf("The HTTP request took %s", elapsed)
	log.Printf("The HTTP response was: %s", response)

	time.Sleep(3 * time.Second)
}

// DoHttpRequest performs an new HTTP request
// that can take between 0 and 500ms to be done
func DoHttpRequest2(result chan<- string, i int) {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(500)
	time.Sleep(time.Duration(n) * time.Millisecond)
	select {
	case result <- "Hellow world!":
		fmt.Printf("Goroutine finished #%d\n", i)
	case <-time.After(1 * time.Second):
		fmt.Printf("Goroutine finished #%d\n", i)
	}
}
