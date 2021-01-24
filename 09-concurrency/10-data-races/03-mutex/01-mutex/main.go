package main

import (
	"fmt"
	"sync"
)

var counter int

func main() {

	gs := 2
	loops := 2

	var wg sync.WaitGroup
	wg.Add(gs)

	var mu sync.Mutex

	for g := 0; g < gs; g++ {
		go func() {
			for i := 0; i < loops; i++ {
				mu.Lock()
				{ // it's a good habit to create artifical code block for better readability
					tmp := counter
					tmp++
					// log.Print("log something") //
					counter = tmp
				}
				mu.Unlock()
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Printf("expected: %d\n", gs*loops)
	fmt.Printf("actual  : %d\n", counter)

}
