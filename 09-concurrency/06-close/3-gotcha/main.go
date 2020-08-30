package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomSet(ch chan int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for _, v := range r.Perm(10) {
		// same as previous example, but every channel sending
		// is executed in it's own go routine
		go func() {
			fmt.Printf("other goroutine: \t send: %v \n", v)
			ch <- v
		}()
	}
	// which means that the above is not a blocking code
	// so the channel gets closed before other goroutines
	// able to finish the sending, panic will occurs
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
