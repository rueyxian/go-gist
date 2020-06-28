package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("hello goroutine: hello")
}
func main() {
	go hello()

	//it's a bad code, you can't assume it will take 1 second to finish other goroutine's operations.
	time.Sleep(1 * time.Second)

	fmt.Println("main goroutine: done")
}
