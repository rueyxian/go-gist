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

	var ch1 chan string
	var ch2 chan string

	go server1(ch1)
	go server2(ch2)

	// It also technically solved nil channel deadlock issue. But it's a very bad design. You shouldn't handle by this way
	select {
	case res := <-ch1:
		fmt.Printf("response: %v \t time: %v \n", res, time.Since(startTime))
	case res := <-ch2:
		fmt.Printf("response: %v \t time: %v \n", res, time.Since(startTime))
	default:
		fmt.Printf("response: - \t time: %v \n", time.Since(startTime))
	}

}
