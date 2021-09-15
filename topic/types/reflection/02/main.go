package main

import (
	"fmt"
	"reflect"
)

type T []interface{ m() }

func (T) m() {}

func main() {
	tp := reflect.TypeOf(new(interface{}))
	tt := reflect.TypeOf(T{})
	fmt.Println(tp, "  ", tt)
	fmt.Println(tp.Kind(), tt.Kind()) // ptr slice
	fmt.Println()

	// Get two interface Types indirectly.
	ti, tim := tp.Elem(), tt.Elem()
	// The next line prints: interface interface
	fmt.Println(ti.Kind(), tim.Kind())
	fmt.Println()

	fmt.Println(tt.Implements(tim))  // true
	fmt.Println(tp.Implements(tim))  // false
	fmt.Println(tim.Implements(tim)) // true
	fmt.Println()

	// All types implement any blank interface type.
	fmt.Println(tp.Implements(ti))  // true
	fmt.Println(tt.Implements(ti))  // true
	fmt.Println(tim.Implements(ti)) // true
	fmt.Println(ti.Implements(ti))  // true
	fmt.Println()
}
