package main

import (
	"fmt"
	"reflect"
)

type T struct{}

func (t T) Sugar(s string) string {
	return fmt.Sprintf("add some sugar into %v", s)
}

func main() {
	rv := reflect.ValueOf(T{})

	rvMethod := rv.MethodByName("Sugar")

	inputs := []reflect.Value{
		reflect.ValueOf("coffee"),
	}

	out := rvMethod.Call(inputs)

	fmt.Printf("%[1]T | %[1]v\n", rv)
	fmt.Printf("%[1]T | %[1]v\n", rvMethod)
	fmt.Printf("%[1]T | %[1]v\n", inputs)
	fmt.Printf("%[1]T | %[1]v\n", out)
}
