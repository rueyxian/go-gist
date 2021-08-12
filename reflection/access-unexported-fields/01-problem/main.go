package main

import (
	"fmt"
	"go-gist/reflection/access-unexported-fields/foo"
)

// ================================================================================

func main() {

	// Let's say we want to access the unexported fields of a struct type – either set or get.
	// We can't simply change the struct's field from unexported to exported.
	// Because it meant to be unexported but we want to access it anyway

	// The code fail to compile

	v := foo.Foo{Exported: "lorem", unexported: 9000}

	fmt.Printf("%[1]T | %[1]v\n", v.Exported)
	fmt.Printf("%[1]T | %[1]v\n", v.unexported)

}

// ================================================================================
