package main

import (
	"fmt"
	"math/rand"
)

func main() {

	ch := make(chan int)

	go func(ch chan int) {
		// blocked by sending, waiting for main goroutine to receive
		ch <- rand.Intn(99)
		fmt.Printf("other goroutine received data from main goroutine: %v \n", <-ch)
	}(ch)
	// blocked by sending, waiting for main goroutine to receive
	ch <- rand.Intn(99)
	fmt.Printf("main goroutine received data from other goroutine: %v \n", <-ch)

}
