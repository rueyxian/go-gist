package main

import (
	"fmt"
	"reflect"
)

func main() {

	rv := reflect.ValueOf([]string{})
	rvs := reflect.MakeSlice(rv.Type(), 100, 1024)

	fmt.Println(rvs.Kind())
	fmt.Println(rvs.Cap())
	fmt.Println(rvs.Len())

}
