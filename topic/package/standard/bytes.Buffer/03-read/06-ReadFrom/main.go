package main

import (
	"bytes"
	"fmt"
	"os"
)

// ReadFrom reads data from r until EOF and appends it to the buffer, growing
// the buffer as needed. The return value n is the number of bytes read. Any
// error except io.EOF encountered during the read is also returned. If the
// buffer becomes too large, ReadFrom will panic with ErrTooLarge.

// func (b *Buffer) ReadFrom(r io.Reader) (n int64, err error) {}

func main() {

	file, _ := os.Open("file.txt")

	buf := bytes.NewBufferString("iddqd ")

	fmt.Println("bytes.Buffer:", buf.String())
	fmt.Println()

	buf.ReadFrom(file)
	fmt.Println("bytes.Buffer:", buf.String())

}
