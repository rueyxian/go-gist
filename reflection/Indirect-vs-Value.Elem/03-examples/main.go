package main

import (
	"fmt"
	"reflect"
)

type T struct{}

func (t T) Foo() {}

type I interface {
	Foo()
}

func main() {

	{
		fmt.Println("if v is a pointer")
		// var v int = 9000
		// rv := reflect.ValueOf(&v)
		rv := reflect.ValueOf(&T{})
		rvi := reflect.Indirect(rv)
		rve := rv.Elem()

		fmt.Printf("\trv : %[1]T | %[1]v | %[2]v\n", rv, rv.Kind())
		fmt.Printf("\trvi: %[1]T | %[1]v | %[2]v\n", rvi, rvi.Kind())
		fmt.Printf("\trve: %[1]T | %[1]v | %[2]v\n", rve, rve.Kind())
		fmt.Println()
	}

	{
		fmt.Println("if v is a nil pointer")
		// var v int
		rv := reflect.ValueOf((*int)(nil))
		rvi := reflect.Indirect(rv)
		rve := rv.Elem()

		fmt.Printf("\trv : %[1]T | %[1]v | %[2]v\n", rv, rv.Kind())
		fmt.Printf("\trvi: %[1]T | %[1]v | %[2]v\n", rvi, rvi.Kind())
		fmt.Printf("\trve: %[1]T | %[1]v | %[2]v\n", rve, rve.Kind())
		fmt.Println()
	}

	{
		fmt.Println("if v is a interface")
		var v I
		rv := reflect.ValueOf(&v)
		rvi := reflect.Indirect(rv)
		rve := rv.Elem()

		fmt.Printf("\trv : %[1]T | %[1]v | %[2]v\n", rv, rv.Kind())
		fmt.Printf("\trvi: %[1]T | %[1]v | %[2]v\n", rvi, rvi.Kind())
		fmt.Printf("\trve: %[1]T | %[1]v | %[2]v\n", rve, rve.Kind())
		fmt.Println()
	}

	{
		fmt.Println("if v is not a pointer or interface")
		rv := reflect.ValueOf(T{})
		rvi := reflect.Indirect(rv)
		// rve := rv.Elem()	// panic

		fmt.Printf("\trv : %[1]T | %[1]v | %[2]v\n", rv, rv.Kind())
		fmt.Printf("\trvi: %[1]T | %[1]v | %[2]v\n", rvi, rvi.Kind())
		fmt.Printf("\trve: can't compile, panic =(\n")
		// fmt.Printf("\trve: %[1]T | %[1]v | %[2]v\n", rve, rve.Kind())
		fmt.Println()
	}

}
