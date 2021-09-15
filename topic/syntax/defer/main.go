package main

import "fmt"

func main() {

	defer A()
	defer B()
	defer C()

}

func A() {
	fmt.Println("A")
}

func B() {
	fmt.Println("B")
}

func C() {
	fmt.Println("C")
}
