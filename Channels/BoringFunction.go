package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan string)
	go boring("boring! ", c)

	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c) // recibe el valor
	}

	fmt.Println("You're boring")
}

func boring(msg string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s %d", msg, i) //la expresión a enviar puede ser cualquier valor adecuado
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
