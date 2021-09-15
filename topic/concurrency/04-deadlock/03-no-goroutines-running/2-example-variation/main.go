package main

import (
	"fmt"
)

func main() {

	ch := make(chan string)

	go func(ch chan string) {
		// blocked by sending, waiting for main goroutine to receive
		ch <- "A"
		fmt.Printf("other goroutine <- main goroutine : %v\n", <-ch)
	}(ch)
	// blocked by sending, waiting for main goroutine to receive
	ch <- "B"
	fmt.Printf("main goroutine <- other goroutine : %v\n", <-ch)

}
