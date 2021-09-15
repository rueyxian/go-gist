package main

import (
	"fmt"
	"reflect"
)

// ================================================================================

type Animal interface {
	MakeSound() string
}

// ================================================================================

type Dog struct{}

func (d *Dog) MakeSound() string {
	return "Bark"
}

// ================================================================================

type Cat struct{}

func (c Cat) MakeSound() string {
	return "Meow"
}

// ================================================================================

func isNil(i interface{}) bool {
	return i == nil || reflect.ValueOf(i).IsNil()
}

func isNilFixed(i interface{}) bool {
	if i == nil {
		return true
	}
	switch reflect.TypeOf(i).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		return reflect.ValueOf(i).IsNil()
	}
	return false
}

// func isNilBetter(i Animal) bool {
//   var ret bool
//   switch i.(type) {
//   case *Dog:
//     v := i.(*Dog)
//     ret = v == nil
//   case Cat:
//     ret = false
//   }

//   return ret
// }

// ================================================================================

func main() {
	var d *Dog = nil
	var a Animal = d

	fmt.Println(isNilFixed(a))

	var c Cat
	a = c
	fmt.Println(isNilFixed(a))

	var m map[string]string
	fmt.Println(isNilFixed(m))

	var s []string
	fmt.Println(isNilFixed(s))

	var ch chan string
	fmt.Println(isNilFixed(ch))
}
