package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// ================================================================================

type t struct {
	foo string
	bar int
}

// ================================================================================

func main() {

	rv := reflect.ValueOf(&t{}).Elem()
	rv0 := rv.Field(0)
	rv1 := rv.Field(1)

	fmt.Println()
	fmt.Printf("%[1]T | %#[1]v\n", unsafe.Pointer(rv0.UnsafeAddr()))
	fmt.Printf("%[1]T | %#[1]v\n", unsafe.Pointer(rv1.UnsafeAddr()))

	// rv0 = reflect.New(rv0.Type()).Elem()
	// rv1 = reflect.New(rv1.Type()).Elem()

	// rv0 = reflect.NewAt(rv0.Type(), unsafe.Pointer(rv0.UnsafeAddr())).Elem()
	// rv1 = reflect.NewAt(rv1.Type(), unsafe.Pointer(rv1.UnsafeAddr())).Elem()
	rv0 = reflect.NewAt(rv0.Type(), unsafe.Pointer(rv0.UnsafeAddr()))
	rv1 = reflect.NewAt(rv1.Type(), unsafe.Pointer(rv1.UnsafeAddr()))

	// fmt.Println()
	// fmt.Printf("%[1]T | %#[1]v\n", rv0)
	// fmt.Printf("%[1]T | %#[1]v\n", rv1)
	// fmt.Println()
	// fmt.Printf("%[1]T | %#[1]v\n", rv0)
	// fmt.Printf("%[1]T | %#[1]v\n", rv1)

	fmt.Println()
	fmt.Printf("%[1]T | %#[1]v\n", unsafe.Pointer(rv0.UnsafeAddr()))
	fmt.Printf("%[1]T | %#[1]v\n", unsafe.Pointer(rv1.UnsafeAddr()))

	rv0 = rv0.Elem()
	rv1 = rv1.Elem()

	fmt.Println()
	fmt.Printf("%[1]T | %#[1]v\n", unsafe.Pointer(rv0.UnsafeAddr()))
	fmt.Printf("%[1]T | %#[1]v\n", unsafe.Pointer(rv1.UnsafeAddr()))

	rv0.Set(reflect.ValueOf(stringPtr("bfg")).Elem())
	rv1.Set(reflect.ValueOf(intPtr(9000)).Elem())

	fmt.Printf("%[1]T | %#[1]v\n", rv)

}

// ================================================================================

func stringPtr(s string) *string { return &s }

func intPtr(i int) *int { return &i }
