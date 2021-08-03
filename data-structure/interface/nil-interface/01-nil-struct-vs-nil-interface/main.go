package main

import "fmt"

type tester interface {
	test()
}

type dummy struct {
}

func (d *dummy) test() {
}

func newNilTester() tester {
	return nil
}

func newNilDummy() *dummy {
	return nil
}

func main() {

	// Under the covers, interfaces are implemented as two elements, a type T and a value V.
	// V is a concrete value such as an int, struct or pointer, never an interface itself, and has type T.
	// For instance, if we store the int value 3 in an interface,
	// the resulting interface value has, schematically, (T=int, V=3).
	// The value V is also known as the interface's dynamic value,
	// since a given interface variable might hold different values V
	// (and corresponding types T) during the execution of the program.

	// An interface value is nil only if the V and T are both unset, (T=nil, V is not set),
	// In particular, a nil interface will always hold a nil type.
	// If we store a nil pointer of type *int inside an interface value,
	// the inner type will be *int regardless of the value of the pointer: (T=*int, V=nil).
	// Such an interface value will therefore be non-nil even when the pointer value V inside is nil.

	// source:
	// https://golang.org/doc/faq#nil_error

	var d1 tester
	d1 = newNilDummy()

	var d2 tester
	d2 = newNilTester()

	var d3 *dummy
	d3 = newNilDummy()

	fmt.Printf("d1: [ T=%[1]T,   V=%[1]v ]\n", d1)
	fmt.Printf("d2: [ T=%[1]T,   V=%[1]v ]\n", d2)
	fmt.Printf("d3: [ T=%[1]T,   V=%[1]v ]\n", d3)

	fmt.Println()

	if d1 != nil {
		fmt.Printf("d1 is not nil\n")
	}

	if d2 != nil {
		fmt.Printf("d2 is not nil\n")
	}

	if d3 != nil {
		fmt.Printf("d3 is not nil\n")
	}

}
