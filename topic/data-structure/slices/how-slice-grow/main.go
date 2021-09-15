package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func a() {
	x := []int{}
	x = append(x, 0)
	x = append(x, 1)
	y := append(x, 2)
	z := append(x, 3)

	monitorSlice(x, "x: ")
	monitorSlice(y, "y: ")
	monitorSlice(z, "z: ")
}

func b() {
	x := []int{}
	x = append(x, 0)
	x = append(x, 1)
	x = append(x, 2)
	y := append(x, 3)
	z := append(x, 4)

	monitorSlice(x, "x: ")
	monitorSlice(y, "y: ")
	monitorSlice(z, "z: ")
}

func main() {
	a()
	fmt.Println("========================================")
	b()
}

func monitorSlice(s []int, discription string) {
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	data := hdr.Data
	fmt.Printf("%s\tdata: %#v\tlen: %v\tcap: %v\t val: %v\n", discription, data, len(s), cap(s), s)
}
