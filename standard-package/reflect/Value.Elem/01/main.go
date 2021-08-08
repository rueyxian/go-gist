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

	// func (v Value) Elem() Value
	// Elem() is a method of relect.Value
	// It simply return reflect.Value of it points to
	// note: reflect.Value does not Implements reflect.Type

	tType := reflect.TypeOf((*reflect.Type)(nil)).Elem()
	tValue := reflect.TypeOf((*reflect.Value)(nil)).Elem()
	fmt.Println(tValue.Implements(tType))
	fmt.Println()

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
	v := reflect.ValueOf(a)
	e := reflect.ValueOf(a).Elem()
	t := reflect.ValueOf(a).Elem().Type()
	// fmt.Printf("%T : %v\n", a, a)
	fmt.Printf("%T : %v\n", v, v)
	fmt.Printf("%T : %v\n", e, e)
	fmt.Printf("%T : %v\n", t, t)
	fmt.Println()
}
