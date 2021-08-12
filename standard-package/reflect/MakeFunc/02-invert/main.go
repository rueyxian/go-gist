package main

import (
	"fmt"
	"reflect"
)

func main() {

	fn := func(args []reflect.Value) []reflect.Value {
		in := args[0]
		out := reflect.MakeSlice(in.Type(), 0, in.Len())
		for i := in.Len() - 1; i >= 0; i-- {
			out = reflect.Append(out, in.Index(i))
		}
		return []reflect.Value{out}
	}

	makeInvert := func(a interface{}) {
		rv := reflect.ValueOf(a).Elem()
		rv.Set(reflect.MakeFunc(rv.Type(), fn))
	}

	//==============================

	var invertInt func([]int) []int
	makeInvert(&invertInt)
	fmt.Println(invertInt([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))

	var invertString func([]string) []string
	makeInvert(&invertString)
	fmt.Println(invertString([]string{"a", "b", "c", "d", "e"}))

}
