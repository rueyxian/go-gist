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

	// Personally, I think this is a bad design, since the default case is for polling the 'waiting for response' status, it should be placed out side the 'select' but within 'for', because the polling is for every cases, not just default case.

	for {
		select {

		case res := <-ch1:
			fmt.Printf("response: %v \t time: %v \n", res, time.Since(startTime))
			return
		case res := <-ch2:
			fmt.Printf("response: %v \t time: %v \n", res, time.Since(startTime))
			return
		default:
			fmt.Printf("waiting for response \t time: %v \n", time.Since(startTime))
			time.Sleep(200 * time.Millisecond)
		}
	}
}
