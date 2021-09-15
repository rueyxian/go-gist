package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var comparer int32

func main() {

	var wg sync.WaitGroup
	comparer = 3

	wg.Add(1)
	go func() {
		swap := atomic.CompareAndSwapInt32(&comparer, 3, 4)
		fmt.Println("swap :", swap)

		fmt.Println("comparer: ", comparer)
		wg.Done()
	}()

	wg.Wait()
}
