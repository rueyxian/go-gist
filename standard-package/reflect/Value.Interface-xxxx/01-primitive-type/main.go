package main

import (
	"fmt"
	"reflect"
)

func main() {

	// func (v Value) Interface() (i interface{})

	var v1 int = 99
	var v2 interface{} = 99
	v3 := reflect.ValueOf(99).Interface()

	fmt.Printf("%T | %v\n", v1, v1)
	fmt.Printf("%T | %v\n", v2, v2)
	fmt.Printf("%T | %v\n", v3, v3)

	fmt.Println()
	fmt.Println(v1 == v2)
	fmt.Println(v1 == v3)
	fmt.Println(v2 == v3)
	fmt.Println()

	// var v4 int
	// v4 = v3

	// ==============================

	// Since int is not the subset of interface{}
	// so assertion is required
	var v4 int = v3.(int)
	_ = v4

	fmt.Println(2 + v3.(int))

}
