package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomSet(ch chan int) {

	for _, v := range rand.Perm(10) {
		fmt.Printf("other goroutine: \t send: %v \n", v)
		ch <- v
	}

	close(ch)
}
func main() {
	rand.Seed(time.Now().UnixNano())

	ch := make(chan int)
	go randomSet(ch)

	for v := range ch {
		fmt.Printf("main goroutine: \t receive: %v \n", v)
	}

}
