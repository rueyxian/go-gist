package main

import (
	"fmt"
	"go-gist/reflection/access-unexported-fields/foo"
	"reflect"
	"unsafe"
)

// ================================================================================

func main() {

	// The solution is to allocate a new memory address, which is a pointer that
	// points to the unexported field.

	v := foo.Foo{Exported: "lorem"}

	rv := reflect.ValueOf(&v).Elem()
	rv0 := rv.Field(0)
	rv1 := rv.Field(1)

	rv0.SetString("bfg")

	rv1 = reflect.NewAt(rv1.Type(), unsafe.Pointer(rv1.UnsafeAddr())).Elem()
	rv1.SetInt(9000)

	fmt.Printf("%[1]T | %#[1]v\n", rv0.Interface())
	fmt.Printf("%[1]T | %#[1]v\n", rv1.Interface())

}

// ================================================================================
