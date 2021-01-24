package main

import (
	"fmt"
	"sync"
)

var counter int

func main() {

	gs := 2
	loops := 5000

	var wg sync.WaitGroup
	wg.Add(gs)

	for g := 0; g < gs; g++ {
		go func() {
			for i := 0; i < loops; i++ {
				counter++
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Printf("expected: %d\n", gs*loops)
	fmt.Printf("actual  : %d\n", counter)

	// even if we use one line of code, say, counter++
	// counter++ itself is not atomic
	// it consist of read, modified, write operation, in assembly code
	// so in hardware level, there is chances that could have preemptive context switch

}
