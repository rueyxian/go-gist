package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	//append is one of the build-in function provides for slices
	//slices does not require to preallocate before using it
	var s []string

	s = append(s, "Hello")
	s = append(s, "world")
	fmt.Println(strings.Join(s, " "))

	//=====================================================

	//bytes.Buffer is another standard library type that is useful zero value
	var b bytes.Buffer
	b.Write([]byte("Hello world"))
	io.Copy(os.Stdout, &b)

}
