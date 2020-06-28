package main

import (
	"fmt"
	"time"
)

func hello(done chan struct{}) {
	fmt.Println("hello goroutine: sleep")
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("hello goroutine: hello")
	done <- struct{}{}
}
func main() {

	// if a channel used solely for blocking, it doesn't matter which type it used for communication between goroutines. Hence, struct{} channel serves the purpose
	done := make(chan struct{})

	fmt.Println("main goroutine: call hello goroutine")
	go hello(done)
	<-done
	fmt.Println("main goroutine: done")
}
