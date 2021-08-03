package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	slice := []int{1, 2, 3, 4, 5}

	slicePtr := unsafe.Pointer(&slice)
	hdr := (*reflect.SliceHeader)(slicePtr)

	dataPtr := unsafe.Pointer(hdr.Data)
	data := *(*[5]int)(unsafe.Pointer(dataPtr))

	fmt.Printf("slice:  %T\n\t%#v\n\tlen: %d\n\tcap: %d\n", slice, slice, len(slice), cap(slice))
	fmt.Println()
	fmt.Printf("&slice: %#v\n", slicePtr)
	fmt.Printf("hdr  : %#v\n", hdr) // &reflect.SliceHeader{Data:0x40e020, Len:4, Cap:4}
	fmt.Println()
	fmt.Printf("&data: %#v\n", dataPtr)
	fmt.Printf("data : %#v\n", data)
	fmt.Println()

}
