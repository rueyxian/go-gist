package main

import "fmt"

func main() {

	f := fib()

	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

}

func fib() func() int {
	a := 0
	b := 1

	return func() int {
		a, b = b, a+b
		return b
	}
}
