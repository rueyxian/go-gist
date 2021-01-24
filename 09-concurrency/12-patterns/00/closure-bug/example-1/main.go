package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"time"
)

// ============================================================

var rnd *rand.Rand

func init() {
	rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func main() {
	closureBug()
	// closureBugFixed()
}

// ============================================================

func closureBug() {
	grs := runtime.NumCPU()
	ch := make(chan string)

	for g := 0; g < grs; g++ {
		go func(idx int) {
			for p := range ch {
				fmt.Printf("[recv_%d]   data: %v\n", idx, p)
			}
		}(g)
	}

	// for g := 0; g < grs; g++ {
	//   go func() {
	//     for p := range ch {
	//       fmt.Printf("[recv_%d]   data: %v\n", g, p)
	//     }
	//   }()
	// }

	works := 20
	for w := 0; w < works; w++ {
		go func() {
			duration := time.Duration(rnd.Intn(250)) * time.Millisecond
			time.Sleep(duration)
			ch <- strconv.Itoa(w)
			fmt.Printf("[send]   data: %v\n", w)
		}()
	}

	time.Sleep(time.Second)
	fmt.Println("========================================")
}

// ============================================================

func closureBugFixed() {
	grs := runtime.NumCPU()
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

	works := 8
	for w := 0; w < works; w++ {
		go func(wrk int) {
			duration := time.Duration(rnd.Intn(250)) * time.Millisecond
			time.Sleep(duration)
			ch <- strconv.Itoa(wrk)
			fmt.Printf("[send]   data: %v\n", wrk)
		}(w)
	}

	time.Sleep(time.Second)
	fmt.Println("========================================")
}
