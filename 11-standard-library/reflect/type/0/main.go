package main

import (
	"fmt"
	"reflect"
)

func main() {

	var i int
	ti := reflect.TypeOf(i)

	fmt.Printf("%v\n", ti)
	fmt.Printf("%T\n", ti)
	fmt.Printf("%T\n", i)

}
