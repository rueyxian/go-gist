package main

import (
	"fmt"
	"math/rand"
	"time"
)

func ping(ch chan int) {
	t := rand.Intn(1000)

	time.Sleep(time.Duration(t) * time.Millisecond)
	time.Sleep(1 * time.Second)
	ch <- t
}

func main() {

	rand.Seed(time.Now().UnixNano())
	ch := make(chan int)
	go ping(ch)
	fmt.Println("start ping")
	fmt.Printf("ping: %v", <-ch)

}
