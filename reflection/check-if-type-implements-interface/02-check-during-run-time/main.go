package main

import (
	"fmt"
	"reflect"
)

type quacker interface {
	quack() string
}

// ================================================================================

type duck struct {
}

func (d duck) quack() string {
	return "quack quack!"
}

// ================================================================================

type chicken struct {
}

// ================================================================================

func main() {
	// ref:
	// https://pkg.go.dev/reflect#TypeOf

	quackerType := reflect.TypeOf((*quacker)(nil)).Elem()

	duckType := reflect.TypeOf((*duck)(nil))
	chickenType := reflect.TypeOf((*chicken)(nil))

	fmt.Println(duckType.Implements(quackerType))
	fmt.Println(chickenType.Implements(quackerType))

	// ==============================
	fmt.Println()

	fmt.Printf("%[1]T | %[1]v\n", reflect.TypeOf((*quacker)(nil)))
	fmt.Printf("%[1]T | %[1]v\n", reflect.TypeOf((*quacker)(nil)).Elem())

}
