package main

import (
	"fmt"
	"reflect"
)

// ================================================================================
type person struct {
	name string
	age  int
}

// ================================================================================

func main() {

	v1 := 99
	v2 := person{"nova", 21}

	var v3 func(int, int) (int, error)
	v4 := func(input string) int {
		return len(input)
	}

	// ==============================

	reflectElem(&v1)
	reflectElem(&v2)
	reflectElem(&v3)
	reflectElem(&v4)

}

// ================================================================================
func reflectElem(a interface{}) {
	v := reflect.ValueOf(a)
	e := reflect.ValueOf(a).Elem()
	t := reflect.ValueOf(a).Elem().Type()
	fmt.Printf("%T : %v\n", a, a)
	fmt.Printf("%T : %v\n", v, v)
	fmt.Printf("%T : %v\n", e, e)
	fmt.Printf("%T : %v\n", t, t)
	fmt.Println()
}
