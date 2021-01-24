package main

import (
	"fmt"
	"math/rand"
)

func main() {

	ch := make(chan int)

	go func(ch chan int) {
		// blocked by receiving, waiting for main goroutine to send
		fmt.Printf("other goroutine <- main goroutine : %v\n", <-ch)
		ch <- rand.Intn(99)
	}(ch)
	// blocked by receiving, waiting for main goroutine to send
	fmt.Printf("main goroutine <- other goroutine", <-ch)
	fmt.Printf("main goroutine <- other goroutine : %v\n", <-ch)
	ch <- rand.Intn(99)

}
