package main

import (
	"bytes"
	"fmt"
)

// ReadBytes reads until the first occurrence of delim in the input,
// returning a slice containing the data up to and including the delimiter.
// If ReadBytes encounters an error before finding a delimiter,
// it returns the data read before the error and the error itself (often io.EOF).
// ReadBytes returns err != nil if and only if the returned data does not end in
// delim.

// func (b *Buffer) ReadBytes(delim byte) (line []byte, err error) {}

func main() {
	buf := bytes.NewBufferString("yoghurt, berries; chapati")

	fmt.Println(buf.String())
	fmt.Println()

	b, _ := buf.ReadBytes(',')
	fmt.Println("bytes.Buffer:", buf.String())
	fmt.Println("[]byte      :", string(b))
	fmt.Println()

	b, _ = buf.ReadBytes(';')
	fmt.Println("bytes.Buffer:", buf.String())
	fmt.Println("[]byte      :", string(b))
	fmt.Println()

}
