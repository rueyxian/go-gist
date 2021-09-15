package main

import (
	"bytes"
	"fmt"
)

// WriteRune appends the UTF-8 encoding of Unicode code point r to the
// buffer, returning its length and an error, which is always nil but is
// included to match bufio.Writer's WriteRune. The buffer is grown as needed;
// if it becomes too large, WriteRune will panic with ErrTooLarge.

// func (b *Buffer) WriteRune(r rune) (n int, err error) {
//     if r < utf8.RuneSelf {
//         b.WriteByte(byte(r))
//         return 1, nil
//     }
//     n = utf8.EncodeRune(b.runeBytes[0:], r)
//     b.Write(b.runeBytes[0:n])
//     return n, nil
// }

// ================================================================================

func main() {

	newRune := rune('喵')
	buf := bytes.NewBufferString("hello")

	fmt.Println(buf.String())

	if _, err := buf.WriteRune(newRune); err != nil {
		panic(err)
	}

	fmt.Println(buf.String())

}
