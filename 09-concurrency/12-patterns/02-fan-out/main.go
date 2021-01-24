package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

var rnd *rand.Rand

func init() {
	rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func main() {
	// fanOut()
	// fanOut_unbuffered_vs_buffered()
	// fanOutV2()
	fanOutSem()

}

func whitespace(n int) string {
	ret := ""
	for i := 0; i < n; i++ {
		ret += " "
	}
	return ret
}

// ================================================================================
func fanOut() {

	// var wg sync.WaitGroup
	startTime := time.Now()
	grs := 1000
	var cSend int64
	var cRecv int64
	ch := make(chan string, grs)

	// wg.Add(grs)

	// ====================
	for g := 0; g < grs; g++ {
		go func(idx int) {
			// defer wg.Done()

			duration := time.Duration(rnd.Intn(250)) * time.Millisecond
			time.Sleep(duration)

			ch <- fmt.Sprintf("result_%v", idx)
			atomic.AddInt64(&cSend, 1)
			padding := whitespace(len(strconv.Itoa(grs)) - len(strconv.Itoa(idx)))
			fmt.Printf("worker_%d%s [send]   duration: %v\n", idx, padding, duration)
		}(g)
	}

	// ====================
	padding := whitespace(len(strconv.Itoa(grs)))
	for i := 0; i < grs; i++ {
		result := <-ch
		atomic.AddInt64(&cRecv, 1)
		fmt.Printf("manager%s [recv]   data: %v\n", padding, result)
	}

	// works := grs
	// for works > 0 {
	//   works--
	//   result := <-ch
	//   fmt.Printf("manager%s [recv]   data: %v\n", padding, result)
	// }

	// ====================
	// wg.Wait()
	fmt.Printf(
		`------------------------------
send count: %d
recv count: %d
time      : %v
==============================
`, cSend, cRecv, time.Since(startTime))
}

// ================================================================================
// Buffered channels help fan out pattern to reduce a little bit of latency
//
// If the channel is unbuffered, all the goroutines (send side) is blocked,
// until the receive side starts signaling, and the goroutines are unblocked one value at a time.
//
// If the channel is buffered, all the goroutines send data to the buffer (if it's blocked)
// and move on without waiting for the receive signal.
//
// To demonstrate that, we put time.Sleep() to give a small pause before send side starts signaling,
// so that we can observe the behavior of data flow of the channel

func fanOut_unbuffered_vs_buffered() {
	fmt.Println("========== unbuffered channel ==========")
	fanOut_unbuffered()
	time.Sleep(time.Second)
	fmt.Println("########################################")
	fmt.Println("========== buffered channel ==========")
	fanOut_buffered()
}

func fanOut_unbuffered() {

	startTime := time.Now()
	grs := 5
	var cSend int64
	var cRecv int64
	ch := make(chan string)

	// ====================
	for g := 0; g < grs; g++ {
		go func(idx int) {
			duration := time.Duration(rnd.Intn(500)) * time.Millisecond
			time.Sleep(duration)

			ch <- fmt.Sprintf("result_%v", idx)
			atomic.AddInt64(&cSend, 1)
			padding := whitespace(len(strconv.Itoa(grs)) - len(strconv.Itoa(idx)))
			fmt.Printf("worker_%d%s [send]   duration: %v\n", idx, padding, duration)
		}(g)
	}

	time.Sleep(time.Second)
	fmt.Println("receive side start signaling...")

	// ====================
	padding := whitespace(len(strconv.Itoa(grs)))
	for i := 0; i < grs; i++ {
		result := <-ch
		atomic.AddInt64(&cRecv, 1)
		fmt.Printf("manager%s [recv]   data: %v\n", padding, result)
	}

	time.Sleep(time.Second)
	// ====================
	fmt.Printf(
		`------------------------------
send count: %d
recv count: %d
time      : %v
==============================
`, cSend, cRecv, time.Since(startTime))
}

func fanOut_buffered() {

	startTime := time.Now()
	grs := 5
	var cSend int64
	var cRecv int64
	ch := make(chan string, grs)

	// ====================
	for g := 0; g < grs; g++ {
		go func(idx int) {
			duration := time.Duration(rnd.Intn(500)) * time.Millisecond
			time.Sleep(duration)

			ch <- fmt.Sprintf("result_%v", idx)
			atomic.AddInt64(&cSend, 1)
			padding := whitespace(len(strconv.Itoa(grs)) - len(strconv.Itoa(idx)))
			fmt.Printf("worker_%d%s [send]   duration: %v\n", idx, padding, duration)
		}(g)
	}

	time.Sleep(time.Second)
	fmt.Println("receive side start signaling...")

	// ====================
	padding := whitespace(len(strconv.Itoa(grs)))
	for i := 0; i < grs; i++ {
		result := <-ch
		atomic.AddInt64(&cRecv, 1)
		fmt.Printf("manager%s [recv]   data: %v\n", padding, result)
	}

	time.Sleep(time.Second)
	// ====================
	fmt.Printf(
		`------------------------------
send count: %d
recv count: %d
time      : %v
==============================
`, cSend, cRecv, time.Since(startTime))
}

// ================================================================================

func fanOutSem() {

	var wg sync.WaitGroup
	var cSend int64
	var cRecv int64
	grs := 100

	ch := make(chan string, grs)
	sem := make(chan struct{}, runtime.NumCPU())

	wg.Add(grs)

	// ====================
	for g := 0; g < grs; g++ {
		go func(idx int) {
			sem <- struct{}{}
			{
				defer wg.Done()

				duration := time.Duration(rnd.Intn(500)) * time.Millisecond
				time.Sleep(duration)

				ch <- fmt.Sprintf("result_%v", idx)
				atomic.AddInt64(&cRecv, 1)
				padding := whitespace(len(strconv.Itoa(grs)) - len(strconv.Itoa(idx)))
				fmt.Printf("worker_%d%s [send]   duration: %v\n", idx, padding, duration)
			}
			<-sem
		}(g)
	}

	// ====================
	padding := whitespace(len(strconv.Itoa(grs)))
	for i := 0; i < grs; i++ {
		result := <-ch
		atomic.AddInt64(&cSend, 1)
		fmt.Printf("manager%s [recv]   data: %v\n", padding, result)
	}

	// ====================
	// wg.Wait()
	fmt.Printf(
		`------------------------------
send count: %d
recv count: %d
==============================
`, cSend, cRecv)

}
