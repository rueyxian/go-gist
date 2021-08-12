package main

import (
	"fmt"
	"reflect"
)

func main() {

	rvf := []reflect.StructField{
		{
			Name: "Name",
			Type: reflect.TypeOf(""),
			Tag:  `json:"Name"`,
		},
		{
			Name: "Price",
			Type: reflect.TypeOf(float64(0)),
			Tag:  `json:"Price"`,
		},
	}

	rt := reflect.StructOf(rvf)
	rv := reflect.New(rt).Elem()

	rv.Field(0).SetString("potato")
	rv.Field(1).SetFloat(0.60)

	fmt.Println(rv)
	fmt.Printf("%#v\n", rv)
	fmt.Println(rv.NumField())

}
