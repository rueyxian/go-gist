package main

import (
	"bytes"
	"fmt"
)

// Write appends the contents of p to the buffer, growing the buffer as
// needed. The return value n is the length of p; err is always nil. If the
// buffer becomes too large, Write will panic with ErrTooLarge.

// func (b *Buffer) Write(p []byte) (n int, err error) {
//   b.lastRead = opInvalid
//   m := b.grow(len(p))
//   return copy(b.buf[m:], p), nil
// }

// ================================================================================

func main() {

	newBytes := []byte("gopher")
	buf := bytes.NewBuffer([]byte("hello"))

	fmt.Println(buf.String())

	v0, err := buf.Write(newBytes)
	if err != nil {
		panic(err)
	}

	fmt.Println(buf.String())

}
