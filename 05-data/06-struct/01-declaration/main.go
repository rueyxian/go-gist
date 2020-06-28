package main

import (
	"fmt"
	"unsafe"
)

type person struct {
	name string
	age  int
}

func main() {

	var a1 person  // zero value
	a2 := person{} // zero value, not recommended, use 'var' for zero value declartion
	a3 := person{"nova", 21}
	a4 := person{"aurora", 17}

	fmt.Println("********** value semantic **********")
	fmt.Printf("    type \t \t T \n")
	fmt.Printf("a1: %T \t %+v \n", a1, a1)
	fmt.Printf("a2: %T \t %+v \n", a2, a2)
	fmt.Printf("a3: %T \t %+v \n", a3, a3)
	fmt.Printf("a4: %T \t %+v \n", a4, a4)

	fmt.Println()
	// ========================================

	// remark: pointer's zero value is nil
	var b1 *person       // zero value (nil)
	var b2 = new(person) // not recommended, use short hand for non-zero value declartion
	b3 := new(person)
	b4 := &person{}
	b5 := &person{"nova", 21}
	// b6 := new(person{"aurora", 17})

	fmt.Println("********** pointer semantic **********")
	fmt.Printf("    type \t \t nil? \t T \t \t *T \n")
	fmt.Printf("b1: %T \t %t \t %v \t %+v \n", b1, b1 == nil, unsafe.Pointer(b1), "")
	// fmt.Printf("b1: %T \t %t \t %v \t %+v \n", b1, b1 == nil, *b1, b1)
	fmt.Printf("b2: %T \t %t \t %v \t %+v \n", b2, b2 == nil, unsafe.Pointer(b2), *b2)
	fmt.Printf("b3: %T \t %t \t %v \t %+v \n", b3, b3 == nil, unsafe.Pointer(b3), *b3)
	fmt.Printf("b4: %T \t %t \t %v \t %+v \n", b4, b4 == nil, unsafe.Pointer(b4), *b4)
	fmt.Printf("b5: %T \t %t \t %v \t %+v \n", b5, b5 == nil, unsafe.Pointer(b5), *b5)
	// fmt.Printf("b6: %T \t %t \t %+v \n", b6, b6 == nil, b6)

	fmt.Println()
	// ========================================

	fmt.Println("********** empty struct **********")
	var es struct{}
	// var ep *struct{}
	fmt.Printf("    type \t  T \t &T \t *T \n")
	fmt.Printf("es: %T \t %v \t %v \t \n", es, es, unsafe.Pointer(&es))
	// fmt.Printf("ep: %T \t %v \t %v \t \n", ep, ep, unsafe.Pointer(&ep))

	fmt.Println()
	// ========================================

}
