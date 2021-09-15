package main

import (
	"fmt"
	"reflect"
)

// ================================================================================

type Sounder interface {
	Sound() string
}

// ================================================================================

type Dog struct {
}

func (d Dog) Sounder() string {
	return "bark"
}

// ================================================================================

type Cat struct {
}

func (c *Cat) Sounder() string {
	return "meow"
}

// ================================================================================

func isNilV1(i interface{}) bool {
	return i == nil || reflect.ValueOf(i).IsNil()
}

func isNilV2(i interface{}) bool {
	if i == nil {
		return true
	}
	switch reflect.TypeOf(i).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		return reflect.ValueOf(i).IsNil()
	}
	return false
}

func isNilV3(i interface{}) bool {
	return i == nil
}

// ================================================================================

func main() {

	// var s Sounder
	var d Dog
	// var c *Cat

	// fmt.Println(d == nil)
	// fmt.Println(d == (*Dog)(nil))

	// fmt.Println(reflect.ValueOf(d).IsNil())
	fmt.Println(isNilV1(d))
	// fmt.Println(reflect.ValueOf(c).IsNil())
	// fmt.Println(c == nil)

	// ==============================

}
