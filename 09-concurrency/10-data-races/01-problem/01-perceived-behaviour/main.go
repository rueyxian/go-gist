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

	for g := 0; g < gs; g++ {
		go func() {
			for i := 0; i < loops; i++ {
				tmp := counter
				tmp++
				counter = tmp
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Printf("expected: %d\n", gs*loops)
	fmt.Printf("actual  : %d\n", counter)

	// if we run the program ...
	// $ go run main.go
	// expected: 4
	// actual  : 4
	//
	// even though program is work as expected, but data races do exists
	// do not get tricked by the perceived behaviour
	// we can check it by running with race detection
	// $ go run -race main.go

}
