package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// ============================================================

func main() {
	select {
	case p := <-ping():
		fmt.Println(p)
	case <-timeout(time.Second):
		fmt.Println("timeout!")
		return
	}

}

// ============================================================

func ping() <-chan string {
	ret := make(chan string)
	go func() {
		duration := time.Duration(rand.Intn(2000)) * time.Millisecond
		time.Sleep(duration)
		ret <- fmt.Sprintf("ping: %v", duration)
	}()
	return ret
}

func timeout(d time.Duration) <-chan struct{} {
	ret := make(chan struct{})
	go func() {
		defer close(ret)
		time.Sleep(d)
		ret <- struct{}{}
	}()
	return ret
}
