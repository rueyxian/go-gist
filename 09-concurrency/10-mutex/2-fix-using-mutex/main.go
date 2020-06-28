package main

import (
	"fmt"
	"sync"
)

var i int

func worker(wg *sync.WaitGroup, mx *sync.Mutex) {
	mx.Lock()
	i++
	mx.Unlock()
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var mx sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go worker(&wg, &mx)
	}

	wg.Wait()

	fmt.Printf("i: %v \n", i)

}
