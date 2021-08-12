package main

import (
	"fmt"
	"reflect"
)

func main() {

	ra := reflect.ArrayOf(9, reflect.TypeOf(0))
	i := reflect.New(ra).Elem().Interface()

	fmt.Printf("%[1]T | %[1]v\n", ra)
	fmt.Printf("%[1]T | %[1]v\n", i)

	// ==============================

	// If we want to use it as an array assertion is required
	//
	// We can actually assert it right away in the above like so:
	// a := reflect.New(ra).Elem().Interface().([9]int)

	a := i.([9]int)
	fmt.Println("len:", len(a))

}

// ================================================================================
