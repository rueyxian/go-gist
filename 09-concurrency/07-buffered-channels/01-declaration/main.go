package main

import (
	"fmt"
)

func main() {
	// buffer in channel is nothing but a temporary memory
	// for the send side when it's blocked

	ch := make(chan string, 2)
	ch <- "hello"
	ch <- "gopher"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
