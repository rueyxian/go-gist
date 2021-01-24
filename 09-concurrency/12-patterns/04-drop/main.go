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
	// drop_base()
	drop()
}

func whitespace(n int) string {
	ret := ""
	for i := 0; i < n; i++ {
		ret += " "
	}
	return ret
}

// ================================================================================
func drop_base() {
	const cap = 2
	ch := make(chan string, cap)

	go func() {
		for p := range ch {
			fmt.Println("employee : recv'd signal :", p)
		}
	}()

	const work = 10
	for w := 0; w < work; w++ {
		select {
		case ch <- "paper":
			fmt.Println("manager : sent signal :", w)
		default:
			fmt.Println("manager : dropped data :", w)
		}
	}

	close(ch)
	fmt.Println("manager : sent shutdown signal")

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------------------")
}

// ================================================================================

func drop() {

	const cap = 10
	const works = 50
	var cSend int64
	var cRecv int64
	var cDrop int64
	ch := make(chan string, cap)
	// ====================
	go func() {

		for task := range ch {
			atomic.AddInt64(&cRecv, 1)
			fmt.Printf("worker  [recv]   data: %v\n", task)
		}
		fmt.Printf("worker  [kill]   \n")
	}()

	// ====================

	for w := 0; w < works; w++ {
		task := fmt.Sprintf("task_%v", w)
		select {
		case ch <- task:
			atomic.AddInt64(&cSend, 1)
			fmt.Printf("manager [send]   data: %v\n", task)
		default:
			atomic.AddInt64(&cDrop, 1)
			fmt.Printf("manager [drop]\n")
		}

	}
	close(ch)
	fmt.Printf("manager [kill]   \n")

	// ====================
	time.Sleep(time.Millisecond * time.Duration(200))
	fmt.Printf(
		`------------------------------
send count: %d
recv count: %d
drop count: %d
==============================
`, cSend, cRecv, cDrop)
}
