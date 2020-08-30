package main

import (
	"fmt"
	"math/rand"
)

func sendData(ch chan<- int) {
	ch <- rand.Intn(99)
	fmt.Printf("sendData goroutine \t type: %T \n", ch)
}

func main() {
	ch := make(chan int)
	go sendData(ch)
	fmt.Printf("main goroutine \t type: %T \n", ch)
	fmt.Println(<-ch)
}
