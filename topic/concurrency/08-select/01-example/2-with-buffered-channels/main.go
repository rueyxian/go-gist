package main

import (
	"fmt"
	"time"
)

var startTime time.Time

func init() {
	startTime = time.Now()
}

func main() {

	ch1 := make(chan string, 2)
	ch2 := make(chan string, 2)

	ch1 <- "apple"
	ch1 <- "orange"
	ch2 <- "plumn"
	ch2 <- "banana"

	time.Sleep(50 * time.Millisecond)
	select {
	case res := <-ch1:
		fmt.Printf("response: %v \t time: %v \n", res, time.Since(startTime))
	case res := <-ch2:
		fmt.Printf("response: %v \t time: %v \n", res, time.Since(startTime))
	}

}
