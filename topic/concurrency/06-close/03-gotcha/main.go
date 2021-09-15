package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomSet(ch chan int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for _, v := range r.Perm(5) {
		// same as previous example, but every channel sending
		// is executed in it's own go routine
		go func(n int) {
			fmt.Printf("other goroutine: \t send: %v \n", n)
			ch <- n
		}(v)
	}
	// every sending execute by individual goroutine, so it's non-blocking code and move on.
	// then channel send close signal so that the channel stop receiving
	// however, the sending run by it's on goroutine, and by chances it still sending data
	// Panic occurs (send on closed channel)

	close(ch)
}

func main() {

	ch := make(chan int)
	go randomSet(ch)

	for v := range ch {
		fmt.Printf("main goroutine: \t receive: %v \n", v)
	}

	// sleep one minute to prevent the program exit right away,
	// because the ch is closed before randomSet goroutine has chances
	// to send even once to main goroutine
	time.Sleep(time.Duration(1) * time.Minute)
}
