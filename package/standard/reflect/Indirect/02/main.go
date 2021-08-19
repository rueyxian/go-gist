package main

import (
	"fmt"
	"reflect"
)

func main() {
	var v1 []int
	var v2 [3]string
	v3 := make(map[int]string)

	rv1 := reflect.ValueOf(&v1)
	rv2 := reflect.ValueOf(&v2)
	rv3 := reflect.ValueOf(&v3)

	rvi1 := reflect.Indirect(rv1)
	rvi2 := reflect.Indirect(rv2)
	rvi3 := reflect.Indirect(rv3)

	fmt.Printf("%v | %v\n", rvi1, rvi1.Kind())
	fmt.Printf("%v | %v\n", rvi2, rvi2.Kind())
	fmt.Printf("%v | %v\n", rvi3, rvi3.Kind())
}
