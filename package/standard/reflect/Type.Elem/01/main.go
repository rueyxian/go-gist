package main

import (
	"fmt"
	"reflect"
)

// ================================================================================

type talker interface {
	talk() string
}

// ================================================================================

type person struct {
	name string
	age  int
}

func (p person) talk() string {
	return fmt.Sprintf("Hello! I'm %v", p.name)
}

// ================================================================================

func main() {

	// func (t *rtype) Elem() Type
	// reflect as a unexported type called rtype which also has a method called Elem()

	// https://cs.opensource.google/go/go/+/refs/tags/go1.16.6:src/reflect/type.go;l=897;bpv=1;bpt=1

	// ==============================

	v1 := 99
	v2 := person{"luna", 21}
	v3 := (talker)(person{"solar", 13})
	var v4 func(int, int) (int, error)
	v5 := func(input string) int {
		return len(input)
	}

	// ==============================

	reflectElem(&v1)
	reflectElem(&v2)
	reflectElem(&v3)
	reflectElem(&v4)
	reflectElem(&v5)

}

// ================================================================================

func reflectElem(a interface{}) {
	v := reflect.TypeOf(a)
	e := reflect.TypeOf(a).Elem()
	// t := reflect.TypeOf(a).Elem().Type()
	// fmt.Printf("%T : %v\n", a, a)
	fmt.Printf("%T : %v\n", v, v)
	fmt.Printf("%T : %v\n", e, e)
	// fmt.Printf("%T : %v\n", t, t)
	fmt.Println()
}
