package main

import (
	"fmt"
	"reflect"
)

func main() {

	ts := reflect.TypeOf("")
	tps := reflect.PtrTo(ts)

	fmt.Printf("%[1]T | %[1]v\n", ts)
	fmt.Printf("%[1]T | %[1]v\n", tps)

}
