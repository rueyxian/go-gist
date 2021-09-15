package main

import (
	"bytes"
	"fmt"
)

// ReadString reads until the first occurrence of delim in the input,
// returning a string containing the data up to and including the delimiter.
// If ReadString encounters an error before finding a delimiter,
// it returns the data read before the error and the error itself (often io.EOF).
// ReadString returns err != nil if and only if the returned data does not end
// in delim.

// func (b *Buffer) ReadString(delim byte) (line string, err error) {}

func main() {
	buf := bytes.NewBufferString("corgi:samoyad#husky")

	fmt.Println(buf.String())
	fmt.Println()

	s, _ := buf.ReadString(':')
	fmt.Println("bytes.Buffer:", buf.String())
	fmt.Println("[]byte      :", s)
	fmt.Println()

	s, _ = buf.ReadString('#')
	fmt.Println("bytes.Buffer:", buf.String())
	fmt.Println("[]byte      :", s)
	fmt.Println()

}
