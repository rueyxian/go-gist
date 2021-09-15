package main

import (
	"bytes"
	"fmt"
)

// Read reads the next len(p) bytes from the buffer or until the buffer
// is drained.  The return value n is the number of bytes read.  If the
// buffer has no data to return, err is io.EOF (unless len(p) is zero);
// otherwise it is nil.
// func (b *Buffer) Read(p []byte) (n int, err error) {}

// ================================================================================

func main() {
	buf := bytes.NewBufferString("1234567890")
	b := make([]byte, 4)

	fmt.Println("bytes.Buffer:", buf.String())
	fmt.Println("[]byte      :", string(b))
	fmt.Println()

	// ==============================
	buf.Read(b)
	fmt.Println("bytes.Buffer:", buf.String())
	fmt.Println("[]byte      :", string(b))
	fmt.Println()

	// ==============================
	buf.Read(b)
	fmt.Println("bytes.Buffer:", buf.String())
	fmt.Println("[]byte      :", string(b))
	fmt.Println()

	// ==============================
	buf.Read(b)
	fmt.Println("bytes.Buffer:", buf.String())
	fmt.Println("[]byte      :", string(b))
	fmt.Println()

}
