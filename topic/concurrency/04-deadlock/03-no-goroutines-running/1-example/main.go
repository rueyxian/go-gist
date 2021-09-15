package main

import (
	"fmt"
)

func main() {

	ch := make(chan string)

	go func(ch chan string) {
		// blocked by receiving, waiting for main goroutine to send
		fmt.Printf("other goroutine <- main goroutine : %v\n", <-ch)
		ch <- "A"
	}(ch)
	// blocked by receiving, waiting for main goroutine to send
	fmt.Printf("main goroutine <- other goroutine", <-ch)
	fmt.Printf("main goroutine <- other goroutine : %v\n", <-ch)
	ch <- "B"

}
