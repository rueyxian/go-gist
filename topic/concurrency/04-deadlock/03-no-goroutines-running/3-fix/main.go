package main

import (
	"fmt"
)

func main() {

	ch := make(chan string)

	go func(ch chan string) {
		fmt.Printf("other goroutine <- main goroutine : %v\n", <-ch)
		ch <- "A"
	}(ch)

	ch <- "B"
	fmt.Printf("main goroutine <- other goroutine : %v\n", <-ch)

}
