package main

import (
	"fmt"
	"reflect"
)

type T struct {
	A string
	B int
}

func main() {

	// 1st way:
	// Passing empty literal into reflect.ValueOf()
	{
		rv := reflect.ValueOf(&T{}).Elem()
		rv.Field(0).SetString("solar")
		rv.Field(1).SetInt(23)
		v := rv.Interface()
		fmt.Println("settable:", rv.CanSet())
		fmt.Printf("%[1]T | %#[1]v\n", v)
		fmt.Println()
	}

	// Passing *T nil makes reflect.Value unsettable.
	// Why? Because it's field(s) are not allocated
	// Consider the code below:
	// var t *T      // t == nil
	// t.A = "solar" // compile error
	// t.B = 21      // compile error
	// Same when working with reflection object.
	{
		rv := reflect.ValueOf((*T)(nil)).Elem()
		fmt.Println("settable:", rv.CanSet())
		fmt.Println()
	}

	// ==============================

	// 2nd way:
	// Passing nil into reflect.TypeOf(), then passing reflect.Type into reflect.New()
	{
		rt := reflect.TypeOf((*T)(nil)).Elem()
		rv := reflect.New(rt).Elem()
		rv.Field(0).SetString("luna")
		rv.Field(1).SetInt(17)
		v := rv.Interface()
		fmt.Println("settable:", rv.CanSet())
		fmt.Printf("%[1]T | %#[1]v\n", v)
		fmt.Println()
	}

	// ==============================

	// So here comes the question: Which is better?

}
