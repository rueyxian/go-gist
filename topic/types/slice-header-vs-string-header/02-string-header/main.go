package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {

	str := "hello"
	hdr := (*reflect.StringHeader)(unsafe.Pointer(&str))
	dataPtr := hdr.Data
	data := (*[1 << 32]byte)(unsafe.Pointer(dataPtr))[:len(str)]

	fmt.Printf("data  \ntype: %T\nval: %#v %q\n\n", data, data, data)

	fmt.Println()

}
