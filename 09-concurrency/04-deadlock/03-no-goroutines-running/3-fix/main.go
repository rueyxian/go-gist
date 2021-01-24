package main

import (
	"fmt"
	"math/rand"
)

func main() {

	ch := make(chan int)

	go func(ch chan int) {
		fmt.Printf("other goroutine <- main goroutine : %v\n", <-ch)
		ch <- rand.Intn(99)
	}(ch)

	ch <- rand.Intn(99)
	fmt.Printf("main goroutine <- other goroutine : %v\n", <-ch)

}
