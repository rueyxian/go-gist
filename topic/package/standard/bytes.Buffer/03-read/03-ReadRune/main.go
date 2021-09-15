package main

import (
	"bytes"
	"fmt"
)

// ReadRune reads and returns the next UTF-8-encoded
// Unicode code point from the buffer.
// If no bytes are available, the error returned is io.EOF.
// If the bytes are an erroneous UTF-8 encoding, it
// consumes one byte and returns U+FFFD, 1.

// func (b *Buffer) ReadRune() (r rune, size int, err error) {}

func main() {

	buf := bytes.NewBufferString("喵😸")

	fmt.Println("bytes.Buffer:", buf.String())
	fmt.Println()

	r, z, _ := buf.ReadRune()
	fmt.Println("bytes.Buffer:", buf.String())
	fmt.Println("rune        :", string(r), " | size: ", z)
	fmt.Println()

	r, z, _ = buf.ReadRune()
	fmt.Println("bytes.Buffer:", buf.String())
	fmt.Println("rune        :", string(r), " | size: ", z)
	fmt.Println()

}
