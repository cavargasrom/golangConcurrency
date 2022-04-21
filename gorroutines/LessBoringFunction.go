package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	go boring2("boring!")

	fmt.Println("Hello 1")
	time.Sleep(2 * time.Second)
	fmt.Println("Hello 2")
}

func boring2(msg string) {

	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}

}
