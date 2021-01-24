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
	runtime.GOMAXPROCS(0)
}

func main() {

	// pooling_base()
	// pooling_fail()
	pooling()
	// poolingV2()

}

func whitespace(n int) string {
	ret := ""
	for i := 0; i < n; i++ {
		ret += " "
	}
	return ret
}

// ================================================================================
// The idea of pooling is we have 'N' number of tasks,
// and distribute those tasks to 'M' number of goroutines
// (If we don't know the optimum value of 'M', take 'M' = number of hardware thread)
// We don't care the sequence task and which goroutine receive which task.
//
// One thing need to be pointed out about this function is that
// when the very last data has been sent and received
// the send side (main goroutine) instantly move on.
// The print function(s) on the receive side might not get called
//
// Because there is a guarentee point on the channel between sending and receiving,
// [ch <- task] and [for task := range ch{ ],
// but can't guarentee if fmt.Printf() (on the receive side) gets called
//
// This code just to demonstrate how a basic pooling pattern looks like.

func pooling_base() {

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

// ================================================================================
// Here is a caveat when using pooling pattern with buffered channel.
// Buffered channel means the guarantee point between sending and receiving is taken away.
// We need something that guarentee that all the sendings will be received before move on, say, wait group.

func pooling_fail() {

	const works = 5000
	var cSend int64
	var cRecv int64
	grs := runtime.NumCPU()
	ch := make(chan string, grs)

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
	}
	close(ch)
	fmt.Printf("manager %s[kill]   \n", padding)

	// ====================
	// time.Sleep(time.Millisecond * time.Duration(500))
	fmt.Printf(
		`------------------------------
send count: %d
recv count: %d
==============================
`, cSend, cRecv)
}

// ================================================================================

func pooling() {

	const works = 4
	var cSend int64
	var cRecv int64
	var wg sync.WaitGroup
	grs := runtime.NumCPU()
	ch := make(chan string)
	wg.Add(grs)

	// ====================
	for g := 0; g < grs; g++ {
		worker := g
		go func() {
			defer wg.Done()

			padding := whitespace(len(strconv.Itoa(grs)) - len(strconv.Itoa(g)))
			for task := range ch {
				atomic.AddInt64(&cRecv, 1)

				// duration := time.Duration(rnd.Intn(10000)) * time.Millisecond
				// duration := time.Second
				// time.Sleep(duration)
				// fmt.Printf("worker_%d%s [recv]   data: %v   duration: %v\n", worker, padding, task, duration)

				fmt.Printf("worker_%d%s [recv]   data: %v\n", worker, padding, task)
			}
			fmt.Printf("worker_%d%s [kill]   \n", worker, padding)
		}()
	}

	// ====================
	padding := whitespace(len(strconv.Itoa(grs)))

	for w := 0; w < works; w++ {
		duration := time.Second
		time.Sleep(duration)

		task := fmt.Sprintf("task_%v", w)
		ch <- task
		atomic.AddInt64(&cSend, 1)
		// fmt.Printf("manager %s[send]   data: %v\n", padding, task)
		fmt.Printf("manager %s[send]   data: %v   duration: %v\n", padding, task, duration)
	}
	close(ch)
	fmt.Printf("manager %s[kill]   \n", padding)

	// ====================
	wg.Wait()
	fmt.Printf(
		`------------------------------
send count: %d
recv count: %d
==============================
`, cSend, cRecv)
}

// ================================================================================

func poolingV2() {

	const works = 4
	var cSend int64
	var cRecv int64
	var wg sync.WaitGroup
	grs := runtime.NumCPU()
	ch := make(chan string, grs)
	wg.Add(grs)

	// ====================
	for g := 0; g < grs; g++ {
		worker := g
		go func() {
			defer wg.Done()

			padding := whitespace(len(strconv.Itoa(grs)) - len(strconv.Itoa(g)))
			for task := range ch {
				atomic.AddInt64(&cRecv, 1)

				// duration := time.Duration(rnd.Intn(10000)) * time.Millisecond
				// duration := time.Second
				// time.Sleep(duration)
				// fmt.Printf("worker_%d%s [recv]   data: %v   duration: %v\n", worker, padding, task, duration)

				fmt.Printf("worker_%d%s [recv]   data: %v\n", worker, padding, task)
			}
			fmt.Printf("worker_%d%s [kill]   \n", worker, padding)
		}()
	}

	// ====================
	padding := whitespace(len(strconv.Itoa(grs)))

	for w := 0; w < works; w++ {
		duration := time.Second
		time.Sleep(duration)

		task := fmt.Sprintf("task_%v", w)
		ch <- task
		atomic.AddInt64(&cSend, 1)
		// fmt.Printf("manager %s[send]   data: %v\n", padding, task)
		fmt.Printf("manager %s[send]   data: %v   duration: %v\n", padding, task, duration)
	}
	close(ch)
	fmt.Printf("manager %s[kill]   \n", padding)

	// ====================
	wg.Wait()
	fmt.Printf(
		`------------------------------
send count: %d
recv count: %d
==============================
`, cSend, cRecv)
}

func poolingV3() {

	const works = 4
	var cSend int64
	var cRecv int64
	var wg sync.WaitGroup
	grs := runtime.NumCPU()
	ch := make(chan string, grs)
	wg.Add(grs)

	// ====================
	for g := 0; g < grs; g++ {
		worker := g
		go func() {
			defer wg.Done()

			padding := whitespace(len(strconv.Itoa(grs)) - len(strconv.Itoa(g)))
			for task := range ch {
				atomic.AddInt64(&cRecv, 1)

				// duration := time.Duration(rnd.Intn(10000)) * time.Millisecond
				// duration := time.Second
				// time.Sleep(duration)
				// fmt.Printf("worker_%d%s [recv]   data: %v   duration: %v\n", worker, padding, task, duration)

				fmt.Printf("worker_%d%s [recv]   data: %v\n", worker, padding, task)
			}
			fmt.Printf("worker_%d%s [kill]   \n", worker, padding)
		}()
	}

	// ====================
	padding := whitespace(len(strconv.Itoa(grs)))

	for w := 0; w < works; w++ {

		duration := time.Second
		time.Sleep(duration)

		task := fmt.Sprintf("task_%v", w)
		ch <- task
		atomic.AddInt64(&cSend, 1)
		// fmt.Printf("manager %s[send]   data: %v\n", padding, task)
		fmt.Printf("manager %s[send]   data: %v   duration: %v\n", padding, task, duration)
	}
	close(ch)
	fmt.Printf("manager %s[kill]   \n", padding)

	// ====================
	wg.Wait()
	fmt.Printf(
		`------------------------------
send count: %d
recv count: %d
==============================
`, cSend, cRecv)
}
