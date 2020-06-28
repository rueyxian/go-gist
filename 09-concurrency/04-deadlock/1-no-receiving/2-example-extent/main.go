package main

import "fmt"

func main() {
	ch := make(chan int)

	// static analysis will stop here. The reason is because the program runs at this point, there is no other goroutines for receiving 'ch'
	ch <- 123

	go func(ch chan int) {
		fmt.Println(<-ch)
	}(ch)

}
