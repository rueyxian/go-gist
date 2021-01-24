package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

var rnd *rand.Rand

func init() {
	rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// ============================================================

func main() {

	fmt.Println("num of goroutines:", runtime.NumGoroutine())
	bug()
	// fix1()
	// fix2()
	fmt.Println("num of goroutines:", runtime.NumGoroutine())
}

// ============================================================

func bug() {
	grs := runtime.NumCPU()
	const works = 20
	ch := make(chan int)

	for gr := 0; gr < grs; gr++ {
		go func() {
			for n := range ch {
				time.Sleep(time.Millisecond * 250)
				fmt.Printf("[worker_%d] %v\n", gr, n)
			}
		}()
	}

	for wrk := 0; wrk < works; wrk++ {
		ch <- wrk
	}

	close(ch)
	time.Sleep(time.Second)
}

// ============================================================
func fix1() {
	grs := runtime.NumCPU()
	const works = 20
	ch := make(chan int)

	for gr := 0; gr < grs; gr++ {
		go func(g int) {
			for n := range ch {
				time.Sleep(time.Millisecond * 250)
				fmt.Printf("[worker_%d] %v\n", g, n)
			}
		}(gr)
	}

	for wrk := 0; wrk < works; wrk++ {
		ch <- wrk
	}

	close(ch)
	time.Sleep(time.Second)
}

// ============================================================
func fix2() {
	grs := runtime.NumCPU()
	const works = 20
	ch := make(chan int)

	for gr := 0; gr < grs; gr++ {
		g := gr
		go func() {
			for n := range ch {
				time.Sleep(time.Millisecond * 250)
				fmt.Printf("[worker_%d] %v\n", g, n)
			}
		}()
	}

	for wrk := 0; wrk < works; wrk++ {
		ch <- wrk
	}

	close(ch)
	time.Sleep(time.Second)
}

// ============================================================
