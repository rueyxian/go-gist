package main

import (
	"fmt"
	"reflect"
)

func main() {
	v := []int{1, 2, 3, 4}
	rv := reflect.ValueOf(&v)
	fmt.Printf("%v | %v\n", rv, rv.Kind())

	rvi := reflect.Indirect(rv)
	fmt.Printf("%v | %v\n", rvi, rvi.Kind())
	fmt.Println()
}
