package main

import (
	"fmt"
	"math/rand"
)

func sendData(ch chan<- int) {
	ch <- rand.Intn(99)
}

func sendAgain(sender chan<- int, receiver <-chan int) {
	sender <- <-receiver
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go sendData(ch1)
	go sendAgain(ch2, ch1)

	// fmt.Printf("ch1: %v ", <-ch1) // this will run into panic coz the data have been received at other goroutine

	fmt.Printf("ch2: %v \n", <-ch2)

}
