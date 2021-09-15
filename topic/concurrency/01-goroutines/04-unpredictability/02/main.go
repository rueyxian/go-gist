package main

import "fmt"

func main() {

	ch := make(chan int)

	const one = 1
	const two = 2

	var cOne int
	var cTwo int

	for i := 0; i < 1000; i++ {

		go func() {
			ch <- one
		}()

		go func() {
			ch <- two
		}()

		if <-ch == 1 {
			cOne++
		} else {
			cTwo++
		}

	}

	fmt.Println("Count One: ", cOne)
	fmt.Println("Count Two: ", cTwo)

}
