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

	var mu sync.RWMutex

	for g := 0; g < gs; g++ {
		go func() {
			for i := 0; i < loops; i++ {
				mu.RLock()
				{
					tmp := counter
					tmp++
					counter = tmp
				}
				mu.RUnlock()
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Printf("expected: %d\n", gs*loops)
	fmt.Printf("actual  : %d\n", counter)

}
