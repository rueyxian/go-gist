package main

import (
	"fmt"
	"reflect"
)

// ================================================================================

func main() {

	var intSwap func(int, int) (int, int)
	makeSwap(&intSwap)
	fmt.Println(intSwap(99, 7))

	var floatSwap func(float64, float64) (float64, float64)
	makeSwap(&floatSwap)
	fmt.Println(floatSwap(3.14, 1.62))

}

// ================================================================================

func swap(in []reflect.Value) []reflect.Value {
	return []reflect.Value{in[1], in[0]}
}

func makeSwap(fptr interface{}) {
	fn := reflect.ValueOf(fptr).Elem()
	v := reflect.MakeFunc(fn.Type(), swap)
	fn.Set(v)
}
