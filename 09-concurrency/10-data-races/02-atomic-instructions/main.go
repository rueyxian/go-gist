package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// var counter int
var counter int32

func main() {

	gs := 2
	loops := 5000

	var wg sync.WaitGroup
	wg.Add(gs)

	for g := 0; g < gs; g++ {
		go func() {
			for i := 0; i < loops; i++ {
				// counter++
				atomic.AddInt32(&counter, 1)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Printf("expected: %d\n", gs*loops)
	fmt.Printf("actual  : %d\n", counter)

}
