package main

import (
	"fmt"
	"math/rand"
	"time"
)

// ============================================================

// var rnd *rand.Rand

func init() {
	// rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
	rand.Seed(time.Now().UTC().UnixNano())
}

// ============================================================

func ping(ch chan<- time.Duration) {
	start := time.Now()
	time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
	ch <- time.Since(start)
}

// ============================================================

func gen() <-chan time.Duration {

}

// ============================================================
func main() {
	primitive()
}

// ============================================================

func primitive() {
	ch := make(chan time.Duration)
	timeout := make(chan time.Duration)
	duration := time.Duration(1000) * time.Millisecond

	// ====================

	go func(d time.Duration) {
		time.Sleep(d)
		timeout <- d
	}(duration)

	// ====================

	go ping(ch)

	// ====================

	for {
		select {
		case ping := <-ch:
			close(timeout)
			fmt.Println("ping: ", ping)
			break
		case timeout := <-timeout:
			fmt.Println("timeout! ", timeout)
			break
		}
	}

	fmt.Println("====================")

}
