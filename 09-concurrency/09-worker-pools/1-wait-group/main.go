package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var startTime time.Time

func init() {
	startTime = time.Now()
}

func process(i int, wg *sync.WaitGroup) {
	fmt.Printf("process %2d  start  %v \n", i, time.Since(startTime))
	time.Sleep(time.Duration(rand.Intn(4000)) * time.Millisecond)
	fmt.Printf("process %2d    end  %v \n", i, time.Since(startTime))

	wg.Done() // wg counter -1, wg.Add(-1) works too

}

func main() {

	fmt.Println("main: start")
	no := 5
	var wg sync.WaitGroup // empty struct wait group

	fmt.Printf("%T \t %+v \n", wg, wg)

	for i := 0; i < no; i++ {
		wg.Add(1) // wg counter +1
		go process(i, &wg)
	}

	// the code blocks here if wg counter is not 0
	// programs will run into panic when:
	// - wg counter is more than 0, and no active goroutines left (deadlock)
	// - wg counter is negative
	wg.Wait()
	fmt.Println("main: end")

}
