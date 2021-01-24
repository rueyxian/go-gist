package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {

	fmt.Println("num of goroutine: ", runtime.NumGoroutine())
	fmt.Println("begin operation...")
	operation()
	time.Sleep(time.Second)
	fmt.Println("end operation...")
	fmt.Println("num of goroutine: ", runtime.NumGoroutine())

}

// Goroutine attempting to send total of 4 value.
// Supposedly the program is expected to receive 4 value correspondently.
// But in reality, it't not always the case.
// By any chance, the stack might ended before the channel get closed
// And the goroutine unable to exit, and not garbage collected.
func operation() {

	ch1 := pipe("A", "B", "C")
	ch2 := pipe("1", "2", "3")

	// ==============================

	output := merge(ch1, ch2)

	// to simulate missing some <-output
	loop := 5
	for i := 0; i < loop; i++ {
		fmt.Println("        recv:", <-output)
	}

	return

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

func merge(chs ...<-chan string) <-chan string {
	ret := make(chan string)
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
		defer close(ret)
		wg.Wait()
	}()

	return ret
}
