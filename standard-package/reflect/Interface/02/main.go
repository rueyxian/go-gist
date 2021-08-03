package main

import (
	"fmt"
	"reflect"
)

// ================================================================================

func main() {

	a1 := reflect.ArrayOf(6, reflect.TypeOf(0))
	a2 := reflect.ArrayOf(9, reflect.TypeOf(0))

	i1 := reflect.New(a1).Interface().(*[6]int)
	i2 := reflect.New(a2).Interface().(*[9]int)

	fmt.Printf("%T  %v\n", i1, i1)
	fmt.Printf("%T  %v\n", i2, i2)

	fmt.Println()

	// ==============================

	fmt.Printf("%T  %v\n", reflect.ValueOf(i1), reflect.ValueOf(i1))
	fmt.Printf("%T  %v\n", reflect.ValueOf(i2), reflect.ValueOf(i2))

	fmt.Println()

	// ==============================

	for i := 0; i < reflect.ValueOf(i1).Elem().Len(); i++ {
		i1[i] = i
	}

	for i := 0; i < 9; i++ {
		i2[i] = i
	}

	fmt.Printf("%T  %v\n", i1, i1)
	fmt.Printf("%T  %v\n", i2, i2)

}
