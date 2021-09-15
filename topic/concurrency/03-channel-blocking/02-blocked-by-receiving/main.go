package main

import (
	"fmt"
)

func hello(done chan bool) {
	fmt.Println("hello goroutine: hello")
	done <- true
}
func main() {

	done := make(chan bool)
	go hello(done)

	// blocking, 'done' in main goroutine as receiving is waiting for 'done' in hello goroutine to send data
	<-done

	fmt.Println("main goroutine: done")
}
