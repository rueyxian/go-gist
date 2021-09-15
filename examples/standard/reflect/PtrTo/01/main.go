package main

import (
	"fmt"
	"reflect"
)

func main() {

	// func PtrTo(t Type) Type
	// basically it's the reverse of
	// func (t *rtype) Elem() Type

	ta := reflect.ArrayOf(5, reflect.TypeOf(int64(0)))

	tpa := reflect.PtrTo(ta)

	fmt.Printf("%[1]T | %[1]v\n", ta)
	fmt.Printf("%[1]T | %[1]v\n", tpa)
	fmt.Printf("%[1]T | %[1]v\n", tpa.Elem())
}
