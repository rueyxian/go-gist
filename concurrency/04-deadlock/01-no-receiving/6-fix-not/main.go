package main

import "fmt"

func main() {
	ch := make(chan int)
	done := make(chan struct{})

	go func(ch chan int, done chan struct{}) {
		fmt.Println(<-ch)
		done <- struct{}{}
	}(ch, done)

	// the code will block here,
	// because main goroutine waiting for receiving struct{}{} first before 123
	// but in other goroutine, it sending 123 first
	<-done
	ch <- 123
}
