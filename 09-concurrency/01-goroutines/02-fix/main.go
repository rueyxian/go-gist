package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("hello goroutine")
}
func main() {
	go hello()
	time.Sleep(1 * time.Second)
	fmt.Println("main goroutine")
}
