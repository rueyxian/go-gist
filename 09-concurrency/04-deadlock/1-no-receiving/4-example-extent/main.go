package main

import "fmt"

func main() {
	ch := make(chan int)

	go func(ch chan int) {
		fmt.Println(<-ch)
	}(ch)

	ch <- 123
	ch <- 456 // panic again, no receiving for this sending

}
