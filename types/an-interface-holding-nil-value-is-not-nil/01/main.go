package main

import "fmt"

func main() {

	var a interface{}
	var b interface{}

	b = (*int)(nil)

	fmt.Printf("a == nil is %t\n", a == nil)
	fmt.Printf("b == nil is %t\n", b == nil)

}
