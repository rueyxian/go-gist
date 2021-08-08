package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {
	// source;
	// https://pkg.go.dev/reflect#TypeOf

	writerType := reflect.TypeOf((*io.Writer)(nil)).Elem()

	fileType := reflect.TypeOf((*os.File)(nil))
	fmt.Println(fileType.Implements(writerType))

}
