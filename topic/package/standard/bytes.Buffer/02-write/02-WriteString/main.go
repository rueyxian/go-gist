package main

import (
	"bytes"
	"fmt"
)

// WriteString appends the contents of s to the buffer, growing the buffer as
// needed. The return value n is the length of s; err is always nil. If the
// buffer becomes too large, WriteString will panic with ErrTooLarge.

// func (b *Buffer) WriteString(s string) (n int, err error) {
//   b.lastRead = opInvalid
//   m := b.grow(len(s))
//   return copy(b.buf[m:], s), nil
// }

// ================================================================================

func main() {
	newString := "gopher"
	buf := bytes.NewBufferString("hello")

	fmt.Println(buf.String())

	buf.WriteString(newString)
	fmt.Println(buf.String())

}
