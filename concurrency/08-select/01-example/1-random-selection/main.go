package main

import (
	"fmt"
	"runtime"
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

	// the reason of making both ch1 and ch2 to be buffered channel of one
	// is because if either one signal first, another one won't be block
	// at it's goroutine, causing memory leak
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)

	go server1(ch1)
	go server2(ch2)

	// if two cases ready at the same time, it will choose randomly
	time.Sleep(50 * time.Millisecond)
	select {
	case res := <-ch1:
		fmt.Printf("response: %v \t time: %v \n", res, time.Since(startTime))
	case res := <-ch2:
		fmt.Printf("response: %v \t time: %v \n", res, time.Since(startTime))
	}

	time.Sleep(50 * time.Millisecond)
	fmt.Println("num of goroutine:", runtime.NumGoroutine())

}
