package main

import (
	"context"
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
	runtime.GOMAXPROCS(0)
}

func main() {

	// cancellation()
	cancellation_bug()
	fmt.Println("num of goroutine:", runtime.NumGoroutine())

}

func whitespace(n int) string {
	ret := ""
	for i := 0; i < n; i++ {
		ret += " "
	}
	return ret
}

// ================================================================================

func cancellation() {
	duration := 150 * time.Millisecond
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	// unbuffered channel instead of buffered
	ch := make(chan string, 1)

	go func() {
		time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)
		ch <- "paper"
	}()

	select {
	case d := <-ch:
		fmt.Println("work complete", d)

	case <-ctx.Done():
		fmt.Println("work cancelled")
	}

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------------------")
}

// ================================================================================

func cancellation_bug() {
	duration := 150 * time.Millisecond
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	// if the channel is unbuffered, and if <-ctx.Done() is signal before d := <-ch,
	// it will moves on, but the other goroutine is blocked since there is no
	// longer a receiver for it, now we have goroutine leak.
	ch := make(chan string)

	go func() {
		time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)
		ch <- "paper"
	}()

	select {
	case d := <-ch:
		fmt.Println("work complete", d)

	case <-ctx.Done():
		fmt.Println("work cancelled")
	}

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------------------")
}
