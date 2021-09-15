package main

import (
	"bytes"
	"fmt"
)

// WriteByte appends the byte c to the buffer, growing the buffer as needed.
// The returned error is always nil, but is included to match bufio.Writer's
// WriteByte. If the buffer becomes too large, WriteByte will panic with
// ErrTooLarge.

// func (b *Buffer) WriteByte(c byte) error {
//     b.lastRead = opInvalid
//     m := b.grow(1)
//     b.buf[m] = c
//     return nil
// }

// ================================================================================

func main() {

	newByte := byte('!')
	buf := bytes.NewBufferString("hello")

	fmt.Println(buf.String())

	if err := buf.WriteByte(newByte); err != nil {
		panic(err)
	}

	fmt.Println(buf.String())

}
