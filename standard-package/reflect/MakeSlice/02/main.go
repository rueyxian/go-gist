package main

import (
	"fmt"
	"reflect"
)

func main() {

	rt1 := reflect.TypeOf([]int{})
	rt2 := reflect.TypeOf(map[string]int{})

	rv1 := reflect.MakeSlice(rt1, 0, 0)
	rv2 := reflect.MakeMap(rt2)

	rv1 = reflect.Append(rv1, reflect.ValueOf(100))

	rv2.SetMapIndex(reflect.ValueOf("bfg"), reflect.ValueOf(9000))

	fmt.Println(rv1)
	fmt.Println(rv2)

}
