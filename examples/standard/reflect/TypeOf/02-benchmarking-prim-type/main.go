package main

import (
	"fmt"
	"reflect"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup
	n := int(1e9)

	wg.Add(1)
	go func() {
		d := testRun(n, typeIntV1)
		fmt.Printf("v1: %v\n", d)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		d := testRun(n, typeIntV2)
		fmt.Printf("v2: %v\n", d)
		wg.Done()
	}()

	wg.Wait()

	// As expected, first way wins: about twice as fast as second way
	// Understandably, second way needs to call an additional step – reflect.Type.Elem() –
	// to get the value that the pointer contains

}

func typeIntV1() reflect.Type {
	return reflect.TypeOf(0)
}

func typeIntV2() reflect.Type {
	return reflect.TypeOf((*int)(nil)).Elem()
}

func testRun(n int, fn func() reflect.Type) time.Duration {
	start := time.Now()
	for i := 0; i < n; i++ {
		fn()
	}
	return time.Now().Sub(start)
}
