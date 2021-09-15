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
}

func main() {

	duration := time.Second

	ctx, cancel := context.WithTimeout(context.Background(), duration)
	ch := gen(ctx)

	fmt.Println("Before: active goroutines", runtime.NumGoroutine())
	fmt.Println("Program exited")

	defer cancel()

	for i := range ch {
		fmt.Println(i)
	}

	fmt.Println("After: active goroutines", runtime.NumGoroutine())
	fmt.Println("Program exited")

}

func gen(ctx context.Context) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for i := 0; ; i++ {
			select {
			case <-time.After(time.Duration(rand.Intn(100)) * time.Millisecond):
				out <- i
			case <-ctx.Done():
				fmt.Println(ctx.Err())
				return
			}
		}
	}()

	return out
}
