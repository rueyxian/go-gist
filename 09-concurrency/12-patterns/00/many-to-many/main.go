package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"time"
)

var rnd *rand.Rand

func init() {
	rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func main() {
	// unbuffered()
	buffered()
}

func unbuffered() {

	grs := runtime.NumCPU()
	// ch := make(chan string, grs)
	ch := make(chan string)

	for g := 0; g < grs; g++ {
		go func(idx int) {

			// duration := time.Duration(rnd.Intn(1000)) * time.Millisecond
			// time.Sleep(duration)
			// data := fmt.Sprintf("data: %2d   duration: %v", g, duration)
			// ch <- data

			for p := range ch {
				fmt.Printf("[recv_%d]   data: %v\n", idx, p)
			}
		}(g)
	}

	works := 4
	for w := 0; w < works; w++ {
		ch <- strconv.Itoa(w)
		fmt.Printf("[send]   data: %v\n", w)
	}

	time.Sleep(time.Second * time.Duration(4))
	fmt.Println("========================================")
}

func buffered() {

	grs := runtime.NumCPU()
	ch := make(chan string, grs)
	// ch := make(chan string)

	for g := 0; g < grs; g++ {
		go func(idx int) {

			for p := range ch {

				duration := time.Duration(rnd.Intn(1000)) * time.Millisecond
				time.Sleep(duration)
				// time.Sleep(time.Second)
				fmt.Printf("[recv_%d]   data: %v   duration: %v\n", idx, p, duration)
			}
		}(g)
	}

	works := 8
	for w := 0; w < works; w++ {
		go func(wrk int) {
			duration := time.Duration(rnd.Intn(1000)) * time.Millisecond
			time.Sleep(duration)
			ch <- strconv.Itoa(wrk)
			fmt.Printf("[send]   data: %v   duration: %v\n", wrk, duration)
		}(w)
	}

	time.Sleep(time.Second * time.Duration(4))
	fmt.Println("========================================")
}
