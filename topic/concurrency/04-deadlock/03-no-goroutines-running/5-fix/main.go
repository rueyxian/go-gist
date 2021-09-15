package main

import (
	"fmt"
)

func sendData(ch1 chan string, ch2 chan string) {

	ch1 <- "A"
	ch2 <- "B"
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go sendData(ch1, ch2)

	fmt.Printf("ch1 : %v \n", <-ch1)
	fmt.Printf("ch2 : %v \n", <-ch2)
}
