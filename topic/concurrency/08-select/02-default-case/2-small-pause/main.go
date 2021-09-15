package main

import (
	"fmt"
	"time"
)

var startTime time.Time

func init() {
	startTime = time.Now()
}

func server1(ch chan string) {
	ch <- "hello, it's-a me, server 1"
}

func server2(ch chan string) {
	ch <- "hello, it's-a me, server 2"
}

func main() {

	ch1 := make(chan string)
	ch2 := make(chan string)

	go server1(ch1)
	go server2(ch2)

	time.Sleep(50 * time.Millisecond)

	// this code is just to prove if given enough of time, other goroutines will get executed, and one of the cases condition can fulfil

	select {
	case res := <-ch1:
		fmt.Printf("response: %v \t time: %v \n", res, time.Since(startTime))
	case res := <-ch2:
		fmt.Printf("response: %v \t time: %v \n", res, time.Since(startTime))
	default:
		fmt.Printf("response: - \t time: %v \n", time.Since(startTime))
	}
}
