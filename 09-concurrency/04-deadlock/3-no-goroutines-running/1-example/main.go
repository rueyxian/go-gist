package main

import (
	"fmt"
	"math/rand"
)

func main() {

	ch := make(chan int)

	go func(ch chan int) {
		// blocked by receiving, waiting for main goroutine to send
		fmt.Printf("other goroutine received data from main goroutine: %v \n", <-ch)
		ch <- rand.Intn(99)
	}(ch)
	// blocked by receiving, waiting for main goroutine to send
	fmt.Printf("main goroutine received data from other goroutine: %v \n", <-ch)
	ch <- rand.Intn(99)

}
