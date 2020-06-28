package main

import (
	"fmt"
)

func hello(done chan bool) {
	fmt.Println("hello goroutine: hello")
	<-done
}

func main() {

	done := make(chan bool)
	go hello(done)

	// blocking, 'done' in main goroutine as sending is waiting for 'done' in hello goroutine to receive data
	done <- true

	fmt.Println("main goroutine: done")
}
