package main

import (
	"fmt"
	"sync"
	"time"
)

// ================================================================================

func main() {

	var wg sync.WaitGroup
	var mx sync.Mutex
	var c1 int
	var c2 int
	n := 100

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			test(mx, &c1, &c2, 1000000)
			wg.Done()
		}()

	}
	wg.Wait()

	fmt.Printf("var decl win count  : %v\n", c1)
	fmt.Printf("short decl win count: %v\n", c2)
	fmt.Printf("draw count          : %v\n", n-c1-c2)

}

// ================================================================================

type test struct {
	wg      sync.WaitGroup
	numCall int
	units   map[string]unit
}

func newTest(unit ...unit) {

}

type unit struct {
	name     string
	fn       func()
	winCount int
}

// ================================================================================

// ================================================================================

func test(mx sync.Mutex, cShort, cVar *int, numCall int) {
	var wg sync.WaitGroup

	wg.Add(1)
	var dShort time.Duration
	go func() {
		dShort = drycall(numCall, callShort)
		// fmt.Printf("short decl: %v\n", d1)
		wg.Done()
	}()

	wg.Wait()

	//short declaration wins
	if dShort < dVar {
		mx.Lock()
		{
			(*cShort)++
		}
		mx.Unlock()
		// fmt.Println("short decl wins")
		return
	}

	// var declaration wins
	if dVar < dShort {
		mx.Lock()
		{
			(*cVar)++
		}
		mx.Unlock()
		// fmt.Println("var decl wins")
		return
	}
}

func callShort() {
	b := byte('?')
	_ = b
}

func callVar() {
	var b byte = '?'
	_ = b
}

func drycall(numCall int, fn func()) time.Duration {
	start := time.Now()
	for i := 0; i < numCall; i++ {
		fn()
	}
	return time.Now().Sub(start)
}
