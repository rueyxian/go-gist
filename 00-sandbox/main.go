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

func whitespace(n int) string {
	ret := ""
	for i := 0; i < n; i++ {
		ret += " "
	}
	return ret
}

func main() {
	// test()
	test0()
	// test1()
	// test2()
	// test3()
}

func test() {

	grs := 1
	_ = grs
	// ch := make(chan string)
	ch := make(chan string, grs)

	fmt.Println(<-ch)
	ch <- "1"
}

func test0() {
	grs := 10
	// ch := make(chan string)
	ch := make(chan string, grs)
	// sem := make(chan struct{}, runtime.NumCPU())
	sem := make(chan struct{}, 2)

	for g := 0; g < grs; g++ {
		go func(idx int) {
			sem <- struct{}{}
			{
				duration := time.Duration(rnd.Intn(1000)) * time.Millisecond
				time.Sleep(duration)

				ch <- fmt.Sprintf("result_%v", idx)
				padding := whitespace(len(strconv.Itoa(grs)) - len(strconv.Itoa(idx)))
				fmt.Printf("worker_%d%s [send]   duration: %v\n", idx, padding, duration)
			}
			<-sem
		}(g)
	}
	time.Sleep(time.Second * time.Duration(5))
	fmt.Println("========== start receiving ==========")

	padding := whitespace(len(strconv.Itoa(grs)))
	for i := 0; i < grs; i++ {
		result := <-ch
		fmt.Printf("manager%s [recv]   data: %v\n", padding, result)
	}

	// time.Sleep(time.Second * time.Duration(5))
	// go func() {
	//   fmt.Println(<-ch)
	// }()

	// ch <- "1"
	// ch <- "2"

	// fmt.Println(<-ch)

}

func test1() {

	grs := 4
	ch := make(chan string, grs)

	for g := 0; g < grs; g++ {
		go func() {

			duration := time.Duration(rnd.Intn(1000)) * time.Millisecond
			time.Sleep(duration)
			data := fmt.Sprintf("data: %2d   duration: %v", g, duration)
			ch <- data
		}()
	}

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	// fmt.Println(<-ch)
	fmt.Println("========================================")
}

func test2() {

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

func test3() {

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
