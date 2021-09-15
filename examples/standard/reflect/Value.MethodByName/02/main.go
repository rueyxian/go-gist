package main

import (
	"fmt"
	"reflect"
)

type T1 struct{}

func (t T1) Foo(b bool) {
	fmt.Println("Foo:", b)
}

type T2 struct{}

func (t T2) Bar(i int, s string) {
	fmt.Println("Bar:", i, s)
}

func invoke(t interface{}, method string, args ...interface{}) {
	inputs := make([]reflect.Value, len(args))

	for i, arg := range args {
		inputs[i] = reflect.ValueOf(arg)
	}

	reflect.ValueOf(t).MethodByName(method).Call(inputs)

}

func main() {

	invoke(T1{}, "Foo", true)
	invoke(T2{}, "Bar", 99, "sweet")

}
