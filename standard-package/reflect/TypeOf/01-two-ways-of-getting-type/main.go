package main

import (
	"fmt"
	"reflect"
)

// ================================================================================

type honker interface {
	honk() string
}

// ================================================================================

type car struct {
}

func (c car) honk() string {
	return "honk honk!"
}

// ================================================================================

func main() {

	// There are two ways of getting reflect.Type (*reflect.rtype) with reflect.TypeOf:
	// 1) passing value as an argument
	// 2) passing pointer (typed nil) as an argument, then getting it's value

	t1 := reflect.TypeOf(0)
	fmt.Printf("%[1]T | %[1]v\n", t1)

	fmt.Println()

	t2 := reflect.TypeOf((*int)(nil)).Elem()
	fmt.Printf("%[1]T | %[1]v\n", t2)

	// ==============================
	fmt.Println()

	// Same when working with composite type
	// However, we often time using 2nd way over 1st way. Why?

	tHonker := reflect.TypeOf((*honker)(nil)).Elem()
	fmt.Printf("%[1]T | %[1]v\n", tHonker)

	tPtrCar := reflect.TypeOf((*car)(nil))
	fmt.Printf("%[1]T | %[1]v\n", tPtrCar)

	tCar := reflect.TypeOf((*car)(nil)).Elem()
	fmt.Printf("%[1]T | %[1]v\n", tCar)

	fmt.Println(tCar.Implements(tHonker))
	fmt.Println(tPtrCar.Implements(tHonker))

}
