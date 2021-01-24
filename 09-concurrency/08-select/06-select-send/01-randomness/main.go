package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		for res := range ch {
			fmt.Printf("[recv]  %v\n", res)
		}

		fmt.Printf("[recv]  kill\n")
	}()

	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond * time.Duration(500))

		select {
		case ch <- "rock":
			fmt.Printf("[send]  count: %v\n", i)
		case ch <- "paper":
			fmt.Printf("[send]  count: %v\n", i)
		case ch <- "scissors":
			fmt.Printf("[send]  count: %v\n", i)
		}
	}

	fmt.Printf("[send]  kill\n")
	close(ch)

}
