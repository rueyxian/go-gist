package main

import (
	"fmt"
	"reflect"
)

// ================================================================================

type Person struct {
	Name string
	Age  int
}

// ================================================================================

func main() {

	{
		rt := reflect.TypeOf((*Person)(nil)).Elem()
		rv := reflect.New(rt).Elem()
		rv.Field(0).SetString("nova")
		rv.Field(1).SetInt(21)
		v := rv.Interface().(Person)

		fmt.Printf("%[1]T | %[1]v\n", rt)
		fmt.Printf("%[1]T | %#[1]v\n", rv)
		fmt.Printf("%[1]T | %#[1]v\n", v)
	}

	{
		rv := reflect.ValueOf(&Person{}).Elem()
		rv.Field(0).SetString("aurora")
		rv.Field(1).SetInt(27)
		v := rv.Interface().(Person)

		fmt.Printf("%[1]T | %#[1]v\n", rv)
		fmt.Printf("%[1]T | %#[1]v\n", v)
	}

}

// ================================================================================
