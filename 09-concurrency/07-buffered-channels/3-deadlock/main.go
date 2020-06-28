package main

import (
	"fmt"
)

func main() {
	ch := make(chan string, 3)
	ch <- "hello"
	ch <- "gopher"
	ch <- "rainbow"
	ch <- "beam" //deadlock
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
