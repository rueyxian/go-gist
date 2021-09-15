package main

import "fmt"

func main() {
	var a []int
	var b []int

	fmt.Println("a == nil : ", a == nil)
	fmt.Println("b == nil : ", b == nil)

	// Go does not support map, slice, function types comparison
	// fmt.Println("a == b   : ", a == b)

}
