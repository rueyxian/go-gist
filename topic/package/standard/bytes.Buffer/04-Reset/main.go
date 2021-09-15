package main

import (
	"bytes"
	"fmt"
)

// Reset resets the buffer so it has no content.
// b.Reset() is the same as b.Truncate(0).

// func (b *Buffer) Reset() { b.Truncate(0) }

func main() {

	buf := bytes.NewBufferString("one does not simply")

	fmt.Println("bytes.Buffer:", buf.String())

	buf.Reset()

	fmt.Println("bytes.Buffer:", buf.String())
}
