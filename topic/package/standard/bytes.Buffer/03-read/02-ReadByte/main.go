package main

import (
	"bytes"
	"fmt"
)

// ReadByte reads and returns the next byte from the buffer.
// If no byte is available, it returns error io.EOF.

// func (b *Buffer) ReadByte() (c byte, err error) {}

func main() {

	buf := bytes.NewBufferString("1234567890")

	b, _ := buf.ReadByte()

	fmt.Println("bytes.Buffer:", buf.String())
	fmt.Println("[]byte      :", string(b))
	fmt.Println()

}

// bufs := bytes.NewBufferString("Learning swift.")
// fmt.Println(bufs.String())
// //读取第一个byte,赋值给b
// b, _ := bufs.ReadByte()
// fmt.Println(bufs.String())
// fmt.Println(string(b))
