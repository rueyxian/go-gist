package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	ch := gen(ctx)
	fmt.Println("Before: active goroutines", runtime.NumGoroutine())

	for i := range ch {
		fmt.Println(i)
		if i >= 5 {
			break
		}
	}

	cancel()
	time.Sleep(time.Millisecond * 500)

	fmt.Println("After: active goroutines", runtime.NumGoroutine())
	fmt.Println("Program exited")

}

func gen(ctx context.Context) <-chan int {
	out := make(chan int)
	count := 1
	go func() {
		defer close(out)
		for {
			select {
			case <-ctx.Done():
				return
			case out <- count:
				count++
			}
		}
	}()
	return out
}
