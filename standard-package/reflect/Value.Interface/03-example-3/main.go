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

	// We don't assert concrete type this time, because
	// fmt package accept interface{} argument.
	// It's unnecessary to assert a concrete type, then pass into fmt method.
	v := rv.Interface()

	fmt.Printf("%T | %v\n", v, v)

}
