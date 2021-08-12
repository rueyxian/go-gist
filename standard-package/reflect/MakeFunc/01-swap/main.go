package main

import (
	"fmt"
	"reflect"
)

func main() {

	swap := func(in []reflect.Value) []reflect.Value {
		return []reflect.Value{in[1], in[0]}
	}

	makeSwap := func(a interface{}) {
		rv := reflect.ValueOf(a).Elem()
		rv.Set(reflect.MakeFunc(rv.Type(), swap))
	}

	var intSwap func(int, int) (int, int)
	makeSwap(&intSwap)
	fmt.Println(intSwap(9, 77))

	var intStringSwap func(int, string) (string, int)
	makeSwap(&intStringSwap)
	fmt.Println(intStringSwap(9000, "bfg"))

}
