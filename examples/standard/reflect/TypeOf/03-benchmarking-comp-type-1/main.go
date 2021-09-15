package main

import (
	"fmt"
	"reflect"
	"sync"
	"time"
)

// ================================================================================

type dumb struct{}

// ================================================================================

func main() {

	var wg sync.WaitGroup
	n := int(1e9)

	wg.Add(1)
	go func() {
		d := testRun(n, typeDumbV1)
		fmt.Printf("v1: %v\n", d)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		d := testRun(n, typeDumbV2)
		fmt.Printf("v2: %v\n", d)
		wg.Done()
	}()

	wg.Wait()

	// Same thing happen when working with composite type:
	// 1st way is faster than 2nd way.

	// But this only valids when that composite type requires minimal
	// memory allocation to instantiate it, e.g. no field struct

}

func typeDumbV1() reflect.Type {
	return reflect.TypeOf(dumb{})
}

func typeDumbV2() reflect.Type {
	return reflect.TypeOf((*dumb)(nil)).Elem()
}

func testRun(n int, fn func() reflect.Type) time.Duration {
	start := time.Now()
	for i := 0; i < n; i++ {
		fn()
	}
	return time.Now().Sub(start)
}
