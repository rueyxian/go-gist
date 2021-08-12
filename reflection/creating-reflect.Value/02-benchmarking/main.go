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
	C bool
}

func main() {

	var wg sync.WaitGroup
	n := int(1e6)

	wg.Add(1)
	go func() {
		fn := func() reflect.Value {
			return reflect.ValueOf(T{})
		}
		fmt.Printf("ValueOf(T{})\n\t%v\n\n", dryRun(n, fn))
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		fn := func() reflect.Value {
			return reflect.ValueOf(&T{})
		}
		fmt.Printf("ValueOf(&T{})\n\t%v\n\n", dryRun(n, fn))
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		fn := func() reflect.Value {
			return reflect.ValueOf((*T)(nil))
		}
		fmt.Printf("ValueOf((*T)(nil))\n\t%v\n\n", dryRun(n, fn))
		wg.Done()
	}()

	// ==============================

	wg.Add(1)
	go func() {
		fn := func() reflect.Value {
			return reflect.New(reflect.TypeOf(T{})).Elem()
		}
		fmt.Printf("New(TypeOf(T{})).Elem()\n\t%v\n\n", dryRun(n, fn))
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		fn := func() reflect.Value {
			return reflect.New(reflect.TypeOf(&T{}).Elem()).Elem()
		}
		fmt.Printf("New(TypeOf(&T{}).Elem()).Elem()\n\t%v\n\n", dryRun(n, fn))
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		fn := func() reflect.Value {
			return reflect.New(reflect.TypeOf((*T)(nil)).Elem()).Elem()
		}
		fmt.Printf("New(TypeOf((*T)(nil)).Elem()).Elem()\n\t%v\n\n", dryRun(n, fn))
		wg.Done()
	}()

	// {
	//   fmt.Println("rt := TypeOf(T{})")
	//   fmt.Println("rv := New(rt).Elem()")

	//   v := "unchanged"
	//   rt := reflect.TypeOf(v)
	//   rv := reflect.New(rt).Elem()
	//   rv.SetString("changed")

	//   fmt.Println("\t", "CanSet():", rv.CanSet())
	//   fmt.Println("\t", v)
	//   fmt.Println()
	// }

	// {
	//   fmt.Println("rt := TypeOf(&T{}).Elem()")
	//   fmt.Println("rv := New(rt).Elem()")

	//   v := "unchanged"
	//   rt := reflect.TypeOf(&v).Elem()
	//   rv := reflect.New(rt).Elem()
	//   rv.SetString("changed")

	//   fmt.Println("\t", "CanSet():", rv.CanSet())
	//   fmt.Println("\t", v)
	//   fmt.Println()
	// }

	// {
	//   fmt.Println("rt := TypeOf((*T)(nil)).Elem()")
	//   fmt.Println("rv := New(rt).Elem()")

	//   var v *string
	//   rt := reflect.TypeOf(v).Elem()
	//   rv := reflect.New(rt).Elem()
	//   rv.SetString("changed")

	//   fmt.Println("\t", "CanSet():", rv.CanSet())
	//   fmt.Println("\t", v)
	//   fmt.Println()
	// }

	// ==============================
	fmt.Println("==============================\n")
	// ==============================

	wg.Wait()

}

func dryRun(n int, fn func() reflect.Value) time.Duration {
	start := time.Now()
	for i := 0; i < n; i++ {
		fn()
	}
	return time.Now().Sub(start)
}
