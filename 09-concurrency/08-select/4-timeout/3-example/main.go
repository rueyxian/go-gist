package main

import (
	"fmt"
	"math/rand"
	"time"
)

var startTime time.Time

func init() {
	startTime = time.Now()
}

func server1(ch chan string) {
	time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)
	ch <- "hello, it's-a me, server 1"
}

func server2(ch chan string) {
	time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)
	ch <- "hello, it's-a me, server 2"
}

func main() {

	rand.Seed(time.Now().UnixNano())
	ch1 := make(chan string)
	ch2 := make(chan string)

	go server1(ch1)
	go server2(ch2)

	// But if you really want to do polling, for loop is required. Declare time channel outside

	pollInt := time.Millisecond
	timeout := time.After(1000 * time.Millisecond)

	for {
		select {
		case res := <-ch1:
			fmt.Printf("response: %v \t time: %v \n", res, time.Since(startTime))
			return
		case res := <-ch2:
			fmt.Printf("response: %v \t time: %v \n", res, time.Since(startTime))
			return
		case <-timeout:
			fmt.Printf("timeout \t time: %v \n", time.Since(startTime))
			return
		default:
		}
		fmt.Printf("waiting for response \t time: %v \n", time.Since(startTime))
		time.Sleep(200 * pollInt)
	}

}
