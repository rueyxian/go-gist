package main

import (
	"fmt"
	"reflect"
)

type Field struct {
	Name  string
	Value interface{}
}

func makeStruct(fields ...Field) interface{} {
	rvfs := make([]reflect.StructField, 0, len(fields))
	for _, v := range fields {
		rvf := reflect.StructField{
			Name: v.Name,
			Type: reflect.TypeOf(v.Value),
		}
		rvfs = append(rvfs, rvf)
	}
	rt := reflect.StructOf(rvfs)
	rv := reflect.New(rt).Elem()

	for i, v := range fields {
		rv.Field(i).Set(reflect.ValueOf(v.Value))
	}

	return rv.Interface()
}

func main() {
	v1 := makeStruct(
		Field{"Name", "potato"},
		Field{"Price", 0.60},
	)
	v2 := makeStruct(
		Field{"Category", "fruits"},
		Field{"Avalaible", true},
		Field{"Items", []string{"blue berry", "avocado", "banana", "durian"}},
	)

	fmt.Printf("%[1]T | %[1]v\n", v1)
	fmt.Printf("%[1]T | %[1]v\n", v2)
}
