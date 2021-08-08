package main

import (
	"fmt"
	"reflect"
)

func main() {

	// create int array type
	rta := reflect.ArrayOf(9, reflect.TypeOf(0))
	ra := reflect.New(rta).Elem()
	a := ra.Interface().([9]int)

	fmt.Printf("%[1]T | %[1]v\n", rta)
	fmt.Printf("%[1]T | %[1]v\n", ra)
	fmt.Printf("%[1]T | %[1]v\n", a)

	fmt.Println()
	// ==============================

	for i := 0; i < len(a); i++ {
		a[i] = i
	}

	fmt.Println(a)

}

// ================================================================================
