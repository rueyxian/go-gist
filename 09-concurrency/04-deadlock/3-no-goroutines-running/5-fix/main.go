package main

import (
	"fmt"
	"math/rand"
)

func sendData(ch1 chan int, ch2 chan int) {

	ch1 <- rand.Intn(99)
	ch2 <- rand.Intn(99)
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go sendData(ch1, ch2)

	fmt.Printf("ch1 : %v \n", <-ch1)
	fmt.Printf("ch2 : %v \n", <-ch2)
}
