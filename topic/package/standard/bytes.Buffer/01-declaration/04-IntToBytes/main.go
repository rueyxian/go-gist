package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {

	fmt.Println(IntToBytes(20))
	fmt.Println(IntToBytes(255))
	fmt.Println(IntToBytes(256))

}

func IntToBytes(n int) []byte {

	buf := bytes.NewBuffer([]byte{})

	binary.Write(buf, binary.BigEndian, int32(n))

	return buf.Bytes()

}
