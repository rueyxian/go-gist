package main

import (
	"context"
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"sync/atomic"
	"time"
)

var rnd *rand.Rand

func init() {
	rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
	runtime.GOMAXPROCS(0)
}

func main() {
	cancellation_base()
}

func whitespace(n int) string {
	ret := ""
	for i := 0; i < n; i++ {
		ret += " "
	}
	return ret
}

// ================================================================================

func cancellation_base() {
	duration := 150 * time.Millisecond
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	ch := make(chan string, 1)

	go func() {
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		ch <- "paper"
	}()

	select {
	case d := <-ch:
		fmt.Println("work complete", d)

	case <-ctx.Done():
		fmt.Println("work cancelled")
	}

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------------------")
}

// ================================================================================

func cancellation() {

	const works = 5000
	var cSend int64
	var cRecv int64
	grs := runtime.NumCPU()
	ch := make(chan string)
	// ====================
	for g := 0; g < grs; g++ {
		worker := g
		go func() {
			padding := whitespace(len(strconv.Itoa(grs)) - len(strconv.Itoa(g)))
			for task := range ch {
				atomic.AddInt64(&cRecv, 1)
				fmt.Printf("worker_%d%s [recv]   data: %v\n", worker, padding, task)
			}
			fmt.Printf("worker_%d%s [kill]   \n", worker, padding)
		}()
	}

	// ====================
	padding := whitespace(len(strconv.Itoa(grs)))

	for w := 0; w < works; w++ {
		task := fmt.Sprintf("task_%v", w)
		ch <- task
		atomic.AddInt64(&cSend, 1)
		fmt.Printf("manager %s[send]   data: %v\n", padding, task)
		// fmt.Printf("manager %s[send]   data: %v   duration: %v\n", padding, task, duration)
	}
	close(ch)
	fmt.Printf("manager %s[kill]   \n", padding)

	// ====================
	// time.Sleep(time.Millisecond * time.Duration(200))
	fmt.Printf(
		`------------------------------
send count: %d
recv count: %d
==============================
`, cSend, cRecv)
}
