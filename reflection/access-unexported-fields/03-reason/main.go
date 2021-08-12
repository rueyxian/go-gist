package main

import (
	"fmt"
	"go-gist/reflection/access-unexported-fields/foo"
	"reflect"
)

// ================================================================================

func main() {

	v := foo.Foo{Exported: "lorem"}

	rv := reflect.ValueOf(&v).Elem()
	rv0 := rv.Field(0)
	rv1 := rv.Field(1)

	// We can check a Value('s) accessibility
	// by using Value.CanSet() & Value.CanInterface() method
	fmt.Println(rv0.CanSet())
	fmt.Println(rv1.CanSet())
	fmt.Println()
	fmt.Println(rv0.CanInterface())
	fmt.Println(rv1.CanInterface())

}

// ================================================================================

func stringPtr(s string) *string { return &s }

func intPtr(i int) *int { return &i }
