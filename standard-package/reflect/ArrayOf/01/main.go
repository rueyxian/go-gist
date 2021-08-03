package main

import (
	"fmt"
	"reflect"
)

type person struct {
	name string
	age  int
}

func main() {

	a1 := reflect.ArrayOf(5, reflect.TypeOf(123))
	a2 := reflect.ArrayOf(9, reflect.TypeOf(true))
	a3 := reflect.ArrayOf(13, reflect.TypeOf(person{}))

	fmt.Printf("%T %v\n", a1, a1)
	fmt.Printf("%T %v\n", a2, a2)
	fmt.Printf("%T %v\n", a3, a3)

}
