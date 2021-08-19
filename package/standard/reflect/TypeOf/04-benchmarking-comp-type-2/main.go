package main

import (
	"fmt"
	"reflect"
	"sync"
	"time"
)

// ================================================================================
type dumb struct {
	fieldA string
	fieldB int
}

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

	// When the struct type contains field(s),
	// the resource of instantiation are exponentially increased
	// Hence, 2nd way are much better

	// Zero value of an struct is empty struct (every field to its default),
	// but zero value of an *struct is nil

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
