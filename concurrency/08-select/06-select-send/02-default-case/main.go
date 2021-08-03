package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

var rnd *rand.Rand

func init() {
	rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func main() {

	var cRecv int64
	var cDrop int64
	ch := make(chan string, 2)

	go func() {
		for res := range ch {

			atomic.AddInt64(&cRecv, 1)
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Printf("[recv]   %v\n", res)
		}
		fmt.Printf("[recv]   kill\n")
	}()

	for i := 0; i < 10; i++ {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		select {
		case ch <- "rock":
			fmt.Printf("[send]   rock\n")
		case ch <- "paper":
			fmt.Printf("[send]   paper\n")
		case ch <- "scissors":
			fmt.Printf("[send]   scissors\n")
		default:

			atomic.AddInt64(&cDrop, 1)
			fmt.Printf("[send]   drop...\n")
		}
	}

	close(ch)
	fmt.Printf("[send]   kill\n")

	time.Sleep(time.Millisecond * time.Duration(250))

	fmt.Printf(
		`------------------------------
drop count: %d
recv count: %d
==============================
`, cDrop, cRecv)

}
