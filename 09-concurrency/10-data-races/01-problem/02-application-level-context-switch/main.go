package main

import (
	"fmt"
	"log"
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
				log.Print("log something") // call log.Print() in between reading and writing
				counter = tmp
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Printf("expected: %d\n", gs*loops)
	fmt.Printf("actual  : %d\n", counter)

	// if	we slot in log function in between reading and writing to force context switch
	// we will see two goroutines accessing the same memory location at the same time

}
