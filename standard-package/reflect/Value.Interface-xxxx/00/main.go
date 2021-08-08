package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func main() {

	//TODO: continue here
	// p := Person{}

	{
		// rv := reflect.ValueOf(&Person{}).Elem()
		rv := reflect.ValueOf((*Person)(nil)).Elem()

		rv.Field(0).SetString("leonardo")
		rv.Field(1).SetInt(43)

		v := rv.Interface().(Person)

		fmt.Printf("%[1]T | %#[1]v\n", rv)
		fmt.Printf("%[1]T | %#[1]v\n", v)
		fmt.Println()
	}

	// ==============================
	{
		rt := reflect.TypeOf((*Person)(nil)).Elem()
		rv := reflect.New(rt).Elem()
		rv.Field(0).SetString("picasso")
		rv.Field(1).SetInt(27)

		// rv := reflect.ValueOf((*person)(nil)).Elem()
		// rv := reflect.ValueOf((*person)(nil))

		fmt.Printf("%[1]T | %[1]v\n", rt)
		fmt.Printf("%[1]T | %#[1]v\n", rv)
		// fmt.Printf("%[1]T | %#[1]v\n", v)

		// fmt.Println()
		// fmt.Println(v1 == v2)
		// fmt.Println(v1 == v3)
		// fmt.Println(v2 == v3)
	}

}
