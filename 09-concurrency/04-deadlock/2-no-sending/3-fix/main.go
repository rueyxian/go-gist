package main

import "fmt"

func main() {
	ch := make(chan int)

	go func(ch chan int) {
		ch <- 9000
	}(ch)

	// unlike previous "no-receiving" example, 'ch' here at the main goroutine act as receiving. 'ch' will block the code, until it gets value from other goroutine, then that value gets printed
	fmt.Println(<-ch)
}
