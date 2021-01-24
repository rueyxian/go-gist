package main

import (
	"fmt"
)

func main() {
	ch := make(chan string, 3)
	ch <- "hello"
	ch <- "gopher"
	ch <- "rainbow"
	ch <- "beam" // the buffer is full, and the code cannot move, hence deadlock
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
