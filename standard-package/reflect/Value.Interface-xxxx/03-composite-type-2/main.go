package main

import (
	"fmt"
	"reflect"
)

func main() {

	rv := reflect.ValueOf(([]string)(nil))

	rv = reflect.Append(rv, reflect.ValueOf("a"))
	rv = reflect.Append(rv, reflect.ValueOf("b"))
	rv = reflect.Append(rv, reflect.ValueOf("c"), reflect.ValueOf("d"))

	// again, assertion is required if we want use it as []string
	v := rv.Interface().([]string)

	fmt.Printf("%T | %v\n", rv, rv)
	fmt.Printf("%T | %v\n", v, v)

}
