package main

import (
	"fmt"
	"runtime"
)

func printNNumber(n int, c chan int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%v \n", <-c)
	}
}

func main() {

	n := 2
	bufSize := 1
	ch := make(chan int, bufSize)

	go printNNumber(n, ch)
	fmt.Printf("NumGoroutine: %v \n", runtime.NumGoroutine())
	ch <- 1
	ch <- 2
	ch <- 3
	// ch <- 4
	// ch <- 5

	fmt.Printf("NumGoroutine: %v \n", runtime.NumGoroutine())
	fmt.Println("main goroutine end")

}
