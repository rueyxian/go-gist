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
	fmt.Println()
	operation()
	time.Sleep(time.Second)
	fmt.Println()
	fmt.Println("end operation...")
	fmt.Println("[after] num of goroutines: ", runtime.NumGoroutine())

}

func operation() {
	done := make(chan struct{}, 2)
	defer close(done)

	ch1 := pipe("A", "B", "C")
	ch2 := pipe("1", "2", "3")

	output := merge(done, ch1, ch2)
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
func pipe(strs ...string) <-chan string {
	ret := make(chan string, len(strs))
	defer close(ret)
	for _, str := range strs {
		ret <- str
	}
	return ret
}

// ============================================================
func merge(done <-chan struct{}, chs ...<-chan string) <-chan string {

	ret := make(chan string)
	var wg sync.WaitGroup
	wg.Add(len(chs))

	output := func(ch <-chan string) {
		defer wg.Done()
		for s := range ch {
			select {
			case ret <- s:
				fmt.Println("send:", s)
			case <-done:
				return
			}
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
