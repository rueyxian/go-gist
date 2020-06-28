package main

import (
	"fmt"
	"math/rand"
)

func sendData(ch chan<- int, done chan struct{}) {
	ch <- rand.Intn(99)
	fmt.Printf("sendData goroutine \t type: %T \n", ch)
	done <- struct{}{}
}

func main() {
	ch := make(chan int)
	done := make(chan struct{})

	go sendData(ch, done)

	fmt.Printf("main goroutine \t type: %T \n", ch)
	<-done

	fmt.Println(<-ch)


}
