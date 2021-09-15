package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

var rnd *rand.Rand

func init() {
	rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func main() {

	fmt.Println("[before] num of goroutines: ", runtime.NumGoroutine())
	fmt.Println()
	operation()
	time.Sleep(time.Second)
	fmt.Println()
	fmt.Println("[after] num of goroutines: ", runtime.NumGoroutine())

}

// func pipe(strs ...string) <-chan string {
//   ret := make(chan string, len(strs))
//   for _, str := range strs {
//     ret <- str
//   }
//   close(ret)
//   return ret
// }

func operation() {

	ch := make(chan time.Duration)
	done := make(chan struct{})

	ping(done, ch)

	fmt.Println("                recv:", <-ch)
	fmt.Println("                recv:", <-ch)
	fmt.Println("                recv:", <-ch)
	fmt.Println("                recv:", <-ch)

	// once the receive completed,
	// send a "done" signal so the goroutine can escape
	// from the select statement.
	//
	// Because if we don't do that, it will stuck at the ch <- duration
	// waiting for receiver.
	// And the goroutine will continue linger until program is terminated (memory leak)
	close(done)

}

func ping(done <-chan struct{}, ch chan<- time.Duration) {

	go func() {
		for {
			duration := time.Duration(rnd.Intn(1000)) * time.Millisecond
			time.Sleep(duration)
			// it's posible to have send & receive signal in a same select statement
			select {
			case ch <- duration: // send signal
				fmt.Println("send:", duration)
			case <-done: // receive signal
				return
			}
		}
	}()
}
