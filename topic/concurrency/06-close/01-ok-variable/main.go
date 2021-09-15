package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomSet(ch chan int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for _, v := range r.Perm(5) {
		fmt.Printf("other goroutine: \t send: %v \n", v)
		ch <- v
	}

	close(ch)
}

func main() {

	ch := make(chan int)
	go randomSet(ch)
	for {
		v, ok := <-ch

		// if ok == false, v = [default value], in this case, 0 is the default value of int
		fmt.Printf("main goroutine: \t receive: %v  %v \n", v, ok)

		if ok == false {
			break
		}

	}

}
