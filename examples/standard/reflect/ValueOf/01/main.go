package main

import (
	"fmt"
	"reflect"
)

func main() {

	a := []int{9, 99}

	v1 := reflect.ValueOf(&a)
	v2 := reflect.ValueOf(a)
	v3 := reflect.ValueOf(&a).Elem()

	fmt.Printf("%T %v\n", v1, v1)
	fmt.Printf("%T %v\n", v2, v2)
	fmt.Printf("%T %v\n", v3, v3)

	fmt.Println()

	// ==============================

	v1 = reflect.Append(v1.Elem(), reflect.ValueOf(3))
	v2 = reflect.Append(v2, reflect.ValueOf(33))
	v3 = reflect.Append(v3, reflect.ValueOf(333))

	// ==============================

	fmt.Printf("%T %v\n", v1, v1)
	fmt.Printf("%T %v\n", v2, v2)
	fmt.Printf("%T %v\n", v3, v3)

}
