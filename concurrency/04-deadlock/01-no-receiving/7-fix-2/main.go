package main

import "fmt"

func main() {
	ch := make(chan int)
	done := make(chan struct{})

	go func(ch chan int, done chan struct{}) {
		fmt.Println(<-ch)
		done <- struct{}{}
	}(ch, done)

	ch <- 123
	<-done
}
