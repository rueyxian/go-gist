package main

import "fmt"

func main() {

	var a *int = nil   // var a *int
	var b *int = nil   // var b *int
	var c *int64 = nil // var c *int64

	fmt.Printf("T=%[1]T | V=%[1]v\n", a)
	fmt.Printf("T=%[1]T | V=%[1]v\n", b)
	fmt.Printf("T=%[1]T | V=%[1]v\n", c)

	fmt.Println("a == b : ", a == b)
	fmt.Println("a == c : ", a == c) // error
	fmt.Println("b == c : ", b == c) // error

}
