package main

import "fmt"

func main() {
	ch := make(chan int)

	// static analysis stops here, because at this point, there is no other goroutines for sending 'ch'
	fmt.Println(<-ch)

	go func(ch chan int) {
		ch <- 9000
	}(ch)

}
