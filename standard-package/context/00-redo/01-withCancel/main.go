package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	after := time.Now().Add(2 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), after)
	defer cancel()

	gen := func(ctx context.Context) <-chan int {
		out := make(chan int)
		go func() {
			defer close(out)
			n := 0
			for {
				select {
				case <-ctx.Done():
					fmt.Println("timeout")
					return
				case out <- n:
					n++
					time.Sleep(500 * time.Millisecond)
				}
			}
		}()
		return out
	}

	for n := range gen(ctx) {
		fmt.Println(n)
	}

	fmt.Println("pause 3 second before program ends")
	time.Sleep(3 * time.Second)

}
