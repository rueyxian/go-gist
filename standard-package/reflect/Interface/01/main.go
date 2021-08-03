package main

import (
	"fmt"
	"reflect"
)

func main() {

	var i1 interface{} = 9
	i2 := reflect.ValueOf(99).Interface()

	fmt.Printf("%T  %v\n", i1, i1)
	fmt.Printf("%T  %v\n", i2, i2)
	fmt.Printf("%T  %v\n", reflect.ValueOf(999), reflect.ValueOf(999))

}
