package main

import (
	"fmt"
	"reflect"
)

func main() {

	// The zero Value represents no value.
	// Its IsValid method returns false, its Kind method returns Invalid,
	// its String method returns "<invalid Value>", and all other methods panic.
	// Most functions and methods never return an invalid value.
	// If one does, its documentation states the conditions explicitly.

	// ref: https://cs.opensource.google/go/go/+/master:src/reflect/value.go;l=26

	var zrv reflect.Value

	fmt.Println(zrv.Kind())
	fmt.Println(zrv.IsValid())
	fmt.Println(zrv)

}
