package main

import "fmt"

func main() {

	ch := make(chan int)

	go func() {
		ch <- 1
	}()

	go func() {
		ch <- 2
	}()

	// question: what is the output?
	fmt.Println(<-ch)

}
