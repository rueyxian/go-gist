package main

import "fmt"

func main() {

	ch1 := twoWayCh()
	fmt.Printf("ch1   type: %T\n", ch1)

	// ====================

	ch2 := sendOnlyCh()
	fmt.Printf("ch2   type: %T\n", ch2)

	// ====================

	ch3 := recvOnlyCh()
	fmt.Printf("ch3   type: %T\n", ch3)
}

func twoWayCh() chan int {
	ret := make(chan int)
	return ret
}

func sendOnlyCh() chan<- int {
	ret := make(chan int)
	return ret
}

func recvOnlyCh() <-chan int {
	ret := make(chan int)
	return ret
}
