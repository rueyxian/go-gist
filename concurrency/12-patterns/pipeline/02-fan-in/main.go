package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {

	ch1 := pipe("A", "B", "C", "D")
	ch2 := pipe("1", "2", "3", "4")

	println("num of goroutines: ", runtime.NumGoroutine())

	// ==============================
	// for i := range merge1(ch1, ch2) {
	//   fmt.Println(i)
	// }

	// ==============================
	// for i := range merge2(ch1, ch2) {
	//   fmt.Println(i)
	// }

	// ==============================
	for i := range merge3(ch1, ch2) {
		fmt.Println(i)
	}

	// ==============================

	time.Sleep(time.Second)

	println("num of goroutines: ", runtime.NumGoroutine())
}

// ============================================================

func pipe(strs ...string) <-chan string {
	ret := make(chan string)
	go func() {
		defer close(ret)
		for _, str := range strs {
			ret <- str
		}
	}()
	return ret
}

// ============================================================

func merge1(chs ...<-chan string) <-chan string {
	ret := make(chan string)
	var wg sync.WaitGroup
	wg.Add(len(chs))

	output := func(ch <-chan string) {
		defer wg.Done()
		for s := range ch {
			ret <- s
		}
	}

	for _, ch := range chs {
		go output(ch)
	}

	go func() {
		defer close(ret)
		wg.Wait()
	}()

	return ret
}

// ============================================================

func merge2(chs ...<-chan string) <-chan string {
	ret := make(chan string)
	var wg sync.WaitGroup
	wg.Add(len(chs))

	output := func(ch <-chan string) {
		defer wg.Done()
		for s := range ch {
			ret <- s
		}
	}

	go func() {
		defer close(ret)
		for _, ch := range chs {
			go output(ch)
		}

		wg.Wait()
	}()

	return ret
}

// ============================================================
// what's different from the above is that
// the output order from the channels does not intersecting
func merge3(chs ...<-chan string) <-chan string {
	ret := make(chan string)
	var wg sync.WaitGroup
	wg.Add(len(chs))

	output := func(ch <-chan string) {
		defer wg.Done()
		for s := range ch {
			ret <- s
		}
	}

	go func() {
		defer close(ret)
		for _, ch := range chs {
			output(ch)
		}

		wg.Wait()
	}()

	return ret
}
