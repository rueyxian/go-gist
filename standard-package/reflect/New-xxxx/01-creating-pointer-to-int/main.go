package main

import (
	"fmt"
	"reflect"
)

func main() {

	rp1 := reflect.New(reflect.TypeOf(0))
	rv1 := rp1.Elem()
	rv1.SetInt(4)

	fmt.Printf("%[1]T | %[1]v\n", rp1)
	fmt.Printf("%[1]T | %[1]v\n", rv1)

	fmt.Println()

	p1 := rp1.Interface().(*int)
	v1 := rv1.Interface()

	fmt.Println(*p1)
	fmt.Println(v1)

	// ==============================
	// basically we are doing something like this ...

	fmt.Println()

	p2 := new(int)
	v2 := 8
	p2 = &v2
	fmt.Println(*p2)
	fmt.Println(v2)

	// ==============================
	// another example: set the value without directly

	fmt.Println()

	rp3 := reflect.New(reflect.TypeOf(0))
	rp3.Elem().SetInt(9)

	p3 := rp3.Interface().(*int)
	fmt.Println(*p3)

	// ==============================
	// basically something like this ...

	fmt.Println()

	p4 := new(int)
	*p4 = 6
	fmt.Println(*p4)

}

// ================================================================================
