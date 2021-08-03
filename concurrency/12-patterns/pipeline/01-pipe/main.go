package main

import (
	"fmt"
)

func main() {

	ch1 := pipe(1, 3, 5)
	ch2 := sqr(ch1)
	for p := range ch2 {
		fmt.Println(p)
	}

	for p := range sqr(pipe(1, 3, 5)) {
		fmt.Println(p)
	}

}

func pipe(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

func sqr(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n * n
		}
	}()
	return out
}
