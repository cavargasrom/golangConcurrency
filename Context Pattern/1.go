package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Doing an HTTP request...")

	start := time.Now()
	response := DoHttpRequest()
	elapsed := time.Since(start)

	log.Printf("The HTTP request took %s", elapsed)
	log.Printf("The HTTP response was: %s", response)
}

// DoHttpRequest performs an new HTTP request
func DoHttpRequest() string {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(500)
	time.Sleep(time.Duration(n) * time.Millisecond)
	return "Hello World"
}
