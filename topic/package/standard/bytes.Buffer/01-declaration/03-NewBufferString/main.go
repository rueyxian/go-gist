package main

import (
	"bytes"
	"fmt"
)

// NewBufferString creates and initializes a new Buffer using string s as its
// initial contents.

//It is intended to prepare a buffer to read an existing string.
//
// In most cases, new(Buffer) (or just declaring a Buffer variable) is
// sufficient to initialize a Buffer.

// func NewBufferString(s string) *Buffer {
//   return &Buffer{buf: []byte(s)}
// }

// ================================================================================

func main() {

	b1 := bytes.NewBufferString("gopher")
	b2 := bytes.NewBuffer([]byte("gopher"))
	b3 := bytes.NewBuffer([]byte{'g', 'o', 'p', 'h', 'e', 'r'})

	fmt.Println("b1:", b1)
	fmt.Println("b2:", b2)
	fmt.Println("b3:", b3)

	// ==============================

	b4 := bytes.NewBufferString("")
	b5 := bytes.NewBuffer([]byte{})
	fmt.Println("b4:", b4)
	fmt.Println("b5:", b5)
}
