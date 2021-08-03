package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {

	fmt.Println("[before] num of goroutines: ", runtime.NumGoroutine())
	fmt.Println("begin operation...")
	operation()
	time.Sleep(time.Second)
	fmt.Println("end operation...")
	fmt.Println("[after] num of goroutines: ", runtime.NumGoroutine())

}

func operation() {

	ch1 := pipe("A", "B", "C")
	ch2 := pipe("1", "2", "3")
	output := merge(ch1, ch2)
	// ==============================

	// for o := range output {
	//   fmt.Println("receive:", o)
	// }

	// to simulate missing some <-output
	loop := 3
	for i := 0; i < loop; i++ {
		fmt.Println("        recv:", <-output)
	}

	return

}

// ============================================================
// make the to-be-returned channel buffer to make it non-blocking
// so that we no need to create a goroutine
func pipe(strs ...string) <-chan string {
	ret := make(chan string, len(strs))
	for _, str := range strs {
		ret <- str
	}
	close(ret)
	return ret
}

// ============================================================

func merge(chs ...<-chan string) <-chan string {
	// To make the channel buffered so that the send operation will
	// (even if there is no receive signal at the moment) put the value
	// into the available buffer and move on.
	//
	// But the question is, how much buffer space do we need?
	// We can't just make the buffer "big enough" to posit that
	// the goroutine won't be blocked.
	// This is not a reliable solution
	ret := make(chan string, 3)
	// ret := make(chan string, 1)

	var wg sync.WaitGroup
	wg.Add(len(chs))

	output := func(ch <-chan string) {
		defer wg.Done()
		for s := range ch {
			ret <- s
			fmt.Println("send:", s)
		}
	}

	for _, ch := range chs {
		go output(ch)
	}

	go func() {
		wg.Wait()
		close(ret)
	}()

	return ret
}

// ============================================================
