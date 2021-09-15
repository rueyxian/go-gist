package main

import (
	"fmt"
	"math/rand"
	"time"
)

var rnd *rand.Rand

func init() {
	rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func main() {

	// if this two channel are unbuffered
	// either one channel get signaled, the other will block
	chA := make(chan string, 0)
	chB := make(chan string, 0)

	go serverA(chA)
	go serverB(chB)

	select {
	case data := <-chA:
		fmt.Println(data)
	case data := <-chB:
		fmt.Println(data)
	}

	time.Sleep(100 * time.Millisecond)

}

func serverA(ch chan string) {
	defer fmt.Println("server A memory released")

	duration := time.Duration(rnd.Intn(1000)) * time.Millisecond
	time.Sleep(duration)
	ch <- fmt.Sprintf("server A : %v", duration)
}

func serverB(ch chan string) {
	defer fmt.Println("server B memory released")

	duration := time.Duration(rnd.Intn(1000)) * time.Millisecond
	time.Sleep(duration)
	ch <- fmt.Sprintf("server B : %v", duration)
}
