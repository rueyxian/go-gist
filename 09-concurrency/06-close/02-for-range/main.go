package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomSet(ch chan int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for _, v := range r.Perm(10) {
		fmt.Printf("other goroutine: \t send: %v \n", v)
		ch <- v
	}

	close(ch)
}
func main() {

	ch := make(chan int)
	go randomSet(ch)

	for v := range ch {
		fmt.Printf("main goroutine: \t receive: %v \n", v)
	}

}
