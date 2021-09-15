package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {

	v1 := (int64)(99)
	rp1 := reflect.NewAt(reflect.TypeOf(int64(0)), unsafe.Pointer(&v1))
	rv1 := reflect.ValueOf(rp1).Pointer()

	fmt.Printf("%[1]T | %[1]v\n", v1)
	fmt.Printf("%[1]T | %[1]v\n", rp1)
	fmt.Printf("%[1]T | %[1]v\n", rv1)

}
