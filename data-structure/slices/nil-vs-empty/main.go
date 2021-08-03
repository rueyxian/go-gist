/*
source:
https://stackoverflow.com/questions/44305170/nil-slices-vs-non-nil-slices-vs-empty-slices-in-go-language
*/

package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

/*
nil and empty slices (with 0 capacity) are not the same, but their observable behavior is the same. By this I mean:

- You can pass them to the builtin len() and cap() functions
- You can for range over them (will be 0 iterations)
- You can slice them (by not violating the restrictions outlined at Spec: Slice expressions; so the result will also be an empty slice)
- Since their length is 0, you can't change their content (appending a value creates a new slice value)
*/

type DummyObject struct {
	value1 string
	value2 int
	value3 bool
}

func main() {

	var s1 []int         // nil slice
	s2 := []int{}        // non-nil, empty slice
	s3 := make([]int, 0) // non-nil, empty slice

	fmt.Println("===== nil slice vs empty slice =====")
	fmt.Printf("s1 | len: %v | cap: %v | s1[:] = %v | s1 == nil  %5v | s1[:] == nil  %5v\n", len(s1), len(s1), s1[:], s1 == nil, s1[:] == nil)
	fmt.Printf("s2 | len: %v | cap: %v | s2[:] = %v | s2 == nil  %5v | s2[:] == nil  %5v\n", len(s2), len(s2), s1[:], s2 == nil, s2[:] == nil)
	fmt.Printf("s3 | len: %v | cap: %v | s3[:] = %v | s3 == nil  %5v | s3[:] == nil  %5v\n", len(s3), len(s3), s1[:], s3 == nil, s3[:] == nil)
	//Note that slicing a nil slice results in a nil slice, slicing a non-nil slice results in a non-nil slice.

	// ============================================================

	// data structure of slices
	/*
		type SliceHeader struct {
					Data uintptr
					Len  int
					Cap  int
			}
	*/

	fmt.Println("")
	fmt.Println("===== data structure =====")

	fmt.Printf("s1 (addr: %p): %+8v\n",
		&s1, *(*reflect.SliceHeader)(unsafe.Pointer(&s1)))
	fmt.Printf("s2 (addr: %p): %+8v\n",
		&s2, *(*reflect.SliceHeader)(unsafe.Pointer(&s2)))
	fmt.Printf("s3 (addr: %p): %+8v\n",
		&s3, *(*reflect.SliceHeader)(unsafe.Pointer(&s3)))

	// ============================================================
	fmt.Println("")
	fmt.Println("===== empty struct type =====")

	var es1 struct{}
	es2 := struct{}{}

	fmt.Printf("es1 addr(hex): %p | addr(dec): %d \n", &es1, (unsafe.Pointer(&es1)))
	fmt.Printf("es2 addr(hex): %p | addr(dec): %d \n", &es2, (unsafe.Pointer(&es2)))

	s2Data := (*(*reflect.SliceHeader)(unsafe.Pointer(&s2))).Data
	s3Data := (*(*reflect.SliceHeader)(unsafe.Pointer(&s3))).Data

	fmt.Printf("s2.data %v \n", s2Data)
	fmt.Printf("s3.data %v \n", s3Data)

	// ============================================================
	fmt.Println("")
	fmt.Println("===== empty slice vs non-empty slice =====")

	s5 := make([]int, 0, 100)  // non-nil, empty slice
	s6 := make([]int, 10, 100) // non-nil, non-empty

	fmt.Printf("s5 | len: %2d | cap: %d |  %v \n", len(s5), cap(s6), s5)
	fmt.Printf("s6 | len: %2d | cap: %d |  %v \n", len(s6), cap(s6), s6)

}
