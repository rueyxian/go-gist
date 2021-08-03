package main

import (
	"fmt"
)

func hello() {
	fmt.Println("hello goroutine")
}
func main() {

	// program is terminated before hello goroutine runs
	go hello()
	fmt.Println("main goroutine")
}
