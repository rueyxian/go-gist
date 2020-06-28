package main

import (
	"fmt"
)

func main() {
	ch := make(chan string, 3)

	fmt.Printf("len: %v \t cap: %v\n", len(ch), cap(ch))

	ch <- "hello"

	fmt.Printf("len: %v \t cap: %v\n", len(ch), cap(ch))

	ch <- "gopher"

	fmt.Printf("len: %v \t cap: %v\n", len(ch), cap(ch))

	<-ch

	fmt.Printf("len: %v \t cap: %v\n", len(ch), cap(ch))

	<-ch

	fmt.Printf("len: %v \t cap: %v\n", len(ch), cap(ch))

	ch <- "something"

	fmt.Printf("len: %v \t cap: %v\n", len(ch), cap(ch))

}
