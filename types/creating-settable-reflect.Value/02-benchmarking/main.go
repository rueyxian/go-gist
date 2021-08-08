package main

import (
	"fmt"
	"reflect"
	"sync"
	"time"
)

type T struct {
	A string
	B int
}

// ================================================================================

func main() {

	var wg sync.WaitGroup
	n := int(1e6)

	wg.Add(1)
	go func() {
		d := testRun(n, versionA)
		fmt.Printf("version A: %v\n", d)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		d := testRun(n, versionB)
		fmt.Printf("version B: %v\n", d)
		wg.Done()
	}()

	wg.Wait()

	// ==============================

	// The answer is pretty obvious really – version A is faster:
	// From interface{} to reflect.Value, only one step is required – reflect.ValueOf()

	// Of course you can go from interface{} to reflect.Type using reflect.Type();
	// Then from reflect.Type() to reflect.Value() using reflect.New().
	// One extra step does not help.

	// The reason I bring this up is because I've seen some code on the internet
	// did version B which boggles me.
	// This experiment for my own confirmation.

}

// ================================================================================

func versionA() reflect.Value {
	rv := reflect.ValueOf(&T{}).Elem()
	rv.Field(0).SetString("iddqd")
	rv.Field(1).SetInt(99)
	return rv
}

func versionB() reflect.Value {
	rt := reflect.TypeOf((*T)(nil)).Elem()
	rv := reflect.New(rt).Elem()
	rv.Field(0).SetString("idkfa")
	rv.Field(1).SetInt(77)
	return rv
}

func testRun(n int, fn func() reflect.Value) time.Duration {
	start := time.Now()
	for i := 0; i < n; i++ {
		fn()
	}
	return time.Now().Sub(start)
}
