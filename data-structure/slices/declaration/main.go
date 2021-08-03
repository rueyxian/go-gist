package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type person struct {
	name string
	age  int
}

func main() {

	fmt.Println("********** value semantic **********")
	var v1 []int
	var v2 = []int{}       // not recommended, use short hand for non-zero value declaration
	var v3 []int = []int{} // horrible!!!
	v4 := []int{}
	v5 := []int{1, 2, 3}
	v6 := make([]int, 3)
	v7 := make([]int, 0, 5)
	v8 := make([]int, 0)

	shv1 := *(*reflect.SliceHeader)(unsafe.Pointer(&v1))
	shv2 := *(*reflect.SliceHeader)(unsafe.Pointer(&v2))
	shv3 := *(*reflect.SliceHeader)(unsafe.Pointer(&v3))
	shv4 := *(*reflect.SliceHeader)(unsafe.Pointer(&v4))
	shv5 := *(*reflect.SliceHeader)(unsafe.Pointer(&v5))
	shv6 := *(*reflect.SliceHeader)(unsafe.Pointer(&v6))
	shv7 := *(*reflect.SliceHeader)(unsafe.Pointer(&v7))
	shv8 := *(*reflect.SliceHeader)(unsafe.Pointer(&v8))

	fmt.Printf("    type \t nil? \t T \t\t sliceHeader \n")
	fmt.Printf("v1: %T \t %t \t %v \t\t %+v \t  \n", v1, v1 == nil, v1, shv1)
	fmt.Printf("v2: %T \t %t \t %v \t\t %+v \t  \n", v2, v2 == nil, v2, shv2)
	fmt.Printf("v3: %T \t %t \t %v \t\t %+v \t  \n", v3, v3 == nil, v3, shv3)
	fmt.Printf("v4: %T \t %t \t %v \t\t %+v \t  \n", v4, v4 == nil, v4, shv4)
	fmt.Printf("v5: %T \t %t \t %v   \t %+v \t  \n", v5, v5 == nil, v5, shv5)
	fmt.Printf("v6: %T \t %t \t %v   \t %+v \t  \n", v6, v6 == nil, v6, shv6)
	fmt.Printf("v7: %T \t %t \t %v \t\t %+v \t  \n", v7, v7 == nil, v7, shv7)
	fmt.Printf("v8: %T \t %t \t %v \t\t %+v \t  \n", v8, v8 == nil, v8, shv8)

	fmt.Println()
	// ========================================

	fmt.Println("********** pointer semantic **********")
	fmt.Printf("    type \t nil? \t T \t\t *T \t  \n")
	var p1 *[]int    // point to nil
	p2 := new([]int) // points to an uninitialized slice with value nil and length 0
	p3 := &[]int{}   // points to an initialized, empty slice with value []int{} and length 0
	p4 := &[]int{1, 2, 3}

	// value of pointer (address of the slice the pointer points to)
	vp1 := unsafe.Pointer(p1)
	vp2 := unsafe.Pointer(p2)
	vp3 := unsafe.Pointer(p3)
	vp4 := unsafe.Pointer(p4)

	//slice header the pointer points to
	shb2 := *(*reflect.SliceHeader)(vp2)
	shb3 := *(*reflect.SliceHeader)(vp3)
	shb4 := *(*reflect.SliceHeader)(vp4)

	fmt.Printf("p1: %T \t %t \t %v\t\t %+v \n", p1, p1 == nil, vp1, "")
	fmt.Printf("p2: %T \t %t \t %v  \t %+v \n", p2, p2 == nil, vp2, shb2)
	fmt.Printf("p3: %T \t %t \t %v  \t %+v \n", p3, p3 == nil, vp3, shb3)
	fmt.Printf("p4: %T \t %t \t %v  \t %+v \n", p4, p4 == nil, vp4, shb4)

	fmt.Println()
	// ========================================

	fmt.Println("********** empty struct **********")
	var es1 struct{}
	es2 := struct{}{}
	fmt.Printf("    type \t  T \t &T(hex) \t &T(dec) \n")
	fmt.Printf("es1: %T \t %v \t %v \t %d \n", es1, es1, unsafe.Pointer(&es1), unsafe.Pointer(&es1))
	fmt.Printf("es2: %T \t %v \t %v \t %d \n", es2, es2, unsafe.Pointer(&es2), unsafe.Pointer(&es2))
	// fmt.Printf("ep: %T \t %v \t %v \t \n", ep, ep, unsafe.Pointer(&ep))

	fmt.Println()
	// ========================================

}
