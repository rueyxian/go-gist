package main

import (
	"fmt"
	"runtime"
	"time"
)

var done = make(chan struct{})

func printNumbers() {
	counter := 1

	for {

		select {
		case <-done:
			return
		default:

			time.Sleep(100 * time.Millisecond)
			counter++
		}

	}

}

func main() {

	go printNumbers()

	fmt.Println("Before: active goroutines", runtime.NumGoroutine())
	time.Sleep(time.Second)

	done <- struct{}{}

	fmt.Println("After: active goroutines", runtime.NumGoroutine())
	fmt.Println("Program exited")

}
