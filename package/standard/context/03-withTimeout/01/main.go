package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

const shortDuration = 1000 * time.Millisecond

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	// Pass a context with a timeout to tell a blocking function that it
	// should abandon its work after the timeout elapses.
	ctx, cancel := context.WithTimeout(context.Background(), shortDuration)
	defer cancel()

	select {
	case d := <-doSomething():
		fmt.Println("done:", d)
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // prints "context deadline exceeded"
	}

}

func doSomething() <-chan time.Duration {
	out := make(chan time.Duration)
	go func() {
		duration := time.Duration(rand.Intn(2000)) * time.Millisecond
		time.Sleep(duration)
		out <- duration
	}()
	return out
}
