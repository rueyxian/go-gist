package main

import (
	"fmt"
	"time"
)

func server1(ch chan string) {
	time.Sleep(500 * time.Millisecond)

	ch <- "hello it's-a me, server 1"
}
func server2(ch chan string) {
	time.Sleep(1000 * time.Millisecond)

	ch <- "hello it's-a me, server 2"
}

var startTime time.Time

func init() {
	startTime = time.Now()
}
func main() {

	ch1 := make(chan string)
	ch2 := make(chan string)
	go server1(ch1)
	go server2(ch2)

	// select statement is blocking without default case. Unblock once one of the case conditions fulfil.

	// it always select case <-ch1, because server2 goroutine is much slower than server1 goroutine.
	select {
	case res := <-ch1:
		fmt.Printf("response: %v \t time: %v \n", res, time.Since(startTime))
	case res := <-ch2:
		fmt.Printf("response: %v \t time: %v \n", res, time.Since(startTime))
	}

}
