package main

import (
	"fmt"
	"time"
)

func main() {

	test()
	// drop()

}

func drop() {
	const cap = 2
	ch := make(chan string, cap)

	go func() {
		for p := range ch {
			fmt.Println("employee : recv'd signal :", p)
		}
	}()

	const work = 10
	for w := 0; w < work; w++ {
		select {
		case ch <- "rock":
			fmt.Println("manager : sent signal :", w)
		case ch <- "paper":
			fmt.Println("manager : sent signal :", w)
		case ch <- "scissors":
			fmt.Println("manager : sent signal :", w)
			// default:
			//   fmt.Println("manager : dropped data :", w)
		}
	}

	close(ch)
	fmt.Println("manager : sent shutdown signal")

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------------------")
}

func test() {

	rand := make(chan int)

	go func() {
		for p := range rand {
			fmt.Println(p)
		}
	}()

	for {
		select {
		case rand <- 0: // no statement
		case rand <- 1:
		}
	}

}
