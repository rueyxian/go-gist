package main

import (
	"fmt"
	"reflect"
)

func main() {

	s1 := reflect.MakeSlice(int, 4, 4)

	fmt.Println(s1)

}
