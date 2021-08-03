package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func randomSet(ch chan int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var wg sync.WaitGroup
	for _, v := range r.Perm(5) {
		wg.Add(1)
		go func(n int) {
			fmt.Printf("other goroutine: \t send: %v \n", n)
			ch <- n
			wg.Done()
		}(v)
	}
	// the solution is to create a wait group to block the code
	// until all the works have been done
	wg.Wait()
	close(ch)
}
func main() {

	ch := make(chan int)

	go randomSet(ch)

	for v := range ch {
		fmt.Printf("main goroutine: \t receive: %v \n", v)
	}

}
