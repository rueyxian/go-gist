package main

import (
	"fmt"
	"reflect"
)

func main() {
	//
	var a *interface{}
	fmt.Println(reflect.TypeOf(a), reflect.ValueOf(a))

	b := (*interface{})(nil)
	fmt.Println(reflect.TypeOf(b), reflect.ValueOf(b))

	var c interface{} = (*interface{})(nil)
	fmt.Println(reflect.TypeOf(c), reflect.ValueOf(c))

	fmt.Println(a == nil, b == nil, c == nil)
}

// source:
// https://stackoverflow.com/questions/43059653/golang-interfacenil-is-nil-or-not
