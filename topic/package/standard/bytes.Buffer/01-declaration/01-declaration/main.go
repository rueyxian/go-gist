package main

import (
	"bytes"
	"fmt"
)

// A Buffer is a variable-sized buffer of bytes with Read and Write methods.
// The zero value for Buffer is an empty buffer ready to use.
type Buffer struct {
	buf      []byte // contents are the bytes buf[off : len(buf)]
	off      int    // read at &buf[off], write at &buf[len(buf)]
	lastRead readOp // last read operation, so that Unread* can work correctly.
}

func main() {

	{
		// Zero value of bytes.Buffer is an empty buffer,
		// and it's ready to use
		var b bytes.Buffer
		fmt.Printf("%T\n", b)

		b.Write([]byte("hello world"))
		fmt.Println(b.String())
		fmt.Println()
	}

	{
		b := &bytes.Buffer{}
		fmt.Printf("%T\n", b)

		b.Write([]byte("hello world"))
		fmt.Println(b.String())
		fmt.Println()
	}

	{ // NewBuffer
		b := bytes.NewBuffer([]byte("hello wolrd"))
		fmt.Printf("%T\n", b)
		fmt.Println(b.String())
		fmt.Println()
	}

	{
		b := bytes.NewBuffer([]byte("hello wolrd"))
		fmt.Printf("%T\n", b)
		fmt.Println(b.String())
		fmt.Println()
	}

}
