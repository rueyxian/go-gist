package main

import (
	"fmt"
	"reflect"
)

func makeStruct(vals ...interface{}) interface{} {
	rvfs := make([]reflect.StructField, 0, len(vals))
	for i, val := range vals {
		rvf := reflect.StructField{
			Name: fmt.Sprintf("F%d", i+1),
			Type: reflect.TypeOf(val),
		}
		rvfs = append(rvfs, rvf)
	}
	rt := reflect.StructOf(rvfs)
	rv := reflect.New(rt).Elem()

	return rv.Interface()
}

func main() {

	v1 := makeStruct("rainbow", true, []int{1, 2, 3})

	rv1 := reflect.ValueOf(v1)

	fmt.Println(v1)

}
