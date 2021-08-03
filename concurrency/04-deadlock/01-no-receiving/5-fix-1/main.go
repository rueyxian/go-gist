package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func(ch chan int) {
		ch <- 123
	}(ch)

	// the code gets blocked by receiving until
	// other goroutine send 123 to main go routine,
	// yes, 123 gets printed
	// but what if we want main goroutine as sender of value 123 for some reason?

	time.Sleep(time.Second * 1)

	fmt.Println(<-ch)
}
