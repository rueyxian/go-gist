package main

import (
	"fmt"
	"go-gist/reflection/access-unexported-fields/foo"
	"reflect"
)

// ================================================================================

func main() {

	// reflect package allows us to examine the fields
	// of a struct, and then access it.

	v := foo.Foo{Exported: "lorem"}

	rv := reflect.ValueOf(&v).Elem()
	rv0 := rv.Field(0)
	rv1 := rv.Field(1)

	// However, this will fail
	// Because Value.Set() method don't allow us to
	// set unexported field directly
	rv0.SetString("bfg")
	rv1.SetInt(9000)

	// Same to Value.Interface()
	fmt.Printf("%[1]T | %#[1]v\n", v)
	fmt.Printf("%[1]T | %#[1]v\n", rv0.Interface())
	fmt.Printf("%[1]T | %#[1]v\n", rv1.Interface())

}

// ================================================================================

func stringPtr(s string) *string { return &s }

func intPtr(i int) *int { return &i }
