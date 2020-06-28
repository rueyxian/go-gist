package main

import "fmt"

func main() {
	ch := make(chan int)

	go func(ch chan int) {
		fmt.Println(<-ch)
	}(ch)

	// no panic but, other goroutine has no chances to print out the receiving value before program is terminated
	ch <- 123
}
