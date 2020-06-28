package main

import (
	"fmt"
	"sync"
)

var i int

func worker(wg *sync.WaitGroup) {
	i++
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go worker(&wg)
	}

	wg.Wait()

	fmt.Printf("i: %v \n", i)

}
