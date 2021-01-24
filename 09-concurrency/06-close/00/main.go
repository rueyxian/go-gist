package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)

	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		ch <- 4
		// close(ch)
	}()

	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)

	go func() {
		fmt.Println(<-ch)
		fmt.Println(<-ch)
		fmt.Println(<-ch)
		fmt.Println(<-ch)
		// close(ch)
	}()

	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)

	time.Sleep(time.Second)

}

func gen() <-chan int {
	out := make(chan int)
	go func() {
		out <- 1
		out <- 2
		out <- 3
		close(out)
	}()
	return out
}
