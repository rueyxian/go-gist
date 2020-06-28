package main

import (
	"fmt"
	"strings"
	"time"
)

func getData(n int, ch chan int) {
	for i := 0; i < n; i++ {
		ch <- i
		lenS := strings.Repeat("▐", len(ch))
		fmt.Printf("   send: %2v  len: %v \n", i, lenS)

		// padding := strings.Repeat(".", (i+1)*2)
		// fmt.Printf("   send: %v %v \n", padding, i)
	}
	close(ch)
}

func main() {

	runtime.NumGoroutine()


	// rand.Seed(time.Now().UnixNano())
	bufCap := 2
	num := 5
	ch := make(chan int, bufCap)

	go getData(num, ch)

	for v := range ch {
		lenS := strings.Repeat("▐", len(ch))
		fmt.Printf("receive: %2v  len: %v \n", v, lenS)

		// padding := strings.Repeat(".", (v+1)*2)
		// fmt.Printf("receive: %v %v \n", padding, v)

		time.Sleep(200 * time.Millisecond)
	}

}
