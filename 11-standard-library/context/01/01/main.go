package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {

	// withoutContext_1()
	withoutContext_2()
	// withContext()
}

// func ex1() {
//   ch := make(chan int)
//   // ====================
//   go func() {
//     for i := 0; i < 5; i++ {
//       ch <- i
//     }
//     close(ch)
//   }()
//   // ====================
//   fmt.Println("num of goroutines: ", runtime.NumGoroutine())
//   for p := range ch {
//     fmt.Println(p)
//   }
//   fmt.Println("num of goroutines: ", runtime.NumGoroutine())
// }

// ============================================================

func withoutContext_1() {
	ch := make(chan int)
	done := make(chan struct{})
	// ====================
	go func() {
		count := 0
		for {
			select {
			case <-done:
				return
			default:
				ch <- count
				count++
			}
		}
	}()
	// ====================
	fmt.Println("num of goroutines: ", runtime.NumGoroutine())
	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond * time.Duration(200))
		fmt.Println(<-ch)
	}
	done <- struct{}{}
	fmt.Println("num of goroutines: ", runtime.NumGoroutine())
}

// ============================================================

func withoutContext_2() {
	ch := make(chan int)
	done := make(chan struct{})
	// ====================
	go func() {
		count := 0
		for {
			select {
			case <-done:
				return
			case ch <- count:
				count++
			}
		}
	}()
	// ====================
	fmt.Println("num of goroutines: ", runtime.NumGoroutine())
	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond * time.Duration(200))
		fmt.Println(<-ch)
	}
	done <- struct{}{}
	fmt.Println("num of goroutines: ", runtime.NumGoroutine())
}

// ============================================================

func withContext() {
	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan int)
	// ====================
	go func() {
		count := 0
		for {
			select {
			case <-ctx.Done():
				return
			case ch <- count:
				count++
			}
		}
	}()
	// ====================
	fmt.Println("num of goroutines: ", runtime.NumGoroutine())
	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond * time.Duration(200))
		fmt.Println(<-ch)
	}
	cancel()

	fmt.Println("num of goroutines: ", runtime.NumGoroutine())
}
