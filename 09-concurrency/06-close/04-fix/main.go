package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func randomSet(ch chan int, done chan struct{}) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var wg sync.WaitGroup
	for _, v := range r.Perm(10) {
		wg.Add(1)
		go func() {
			fmt.Printf("other goroutine: \t send: %v \n", v)
			ch <- v
			wg.Done()
		}()
	}
	// the solution is to create a wait group to block the code
	// until all the works have been done
	wg.Wait()
	close(ch)
	done <- struct{}{}
}
func main() {

	done := make(chan struct{})
	ch := make(chan int)

	go randomSet(ch, done)

	for v := range ch {
		fmt.Printf("main goroutine: \t receive: %v \n", v)
	}

	<-done
}
