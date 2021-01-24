package main

import (
	"fmt"
	"reflect"
)

func main() {
	d1 := []string{"one", "two", "three"}
	d2 := []int{1, 2, 3}
	d3 := "gopher"

	test(d1)
	test(d2)
	test(d3)

	fmt.Printf("%+v  %+v  %+v\n", reflect.TypeOf(d1).Kind(), reflect.TypeOf(d1), reflect.ValueOf(d1))
	fmt.Printf("%+v  %+v  %+v\n", reflect.TypeOf(d2).Kind(), reflect.TypeOf(d2), reflect.ValueOf(d2))
	fmt.Printf("%+v  %+v  %+v\n", reflect.TypeOf(d3).Kind(), reflect.TypeOf(d3), reflect.ValueOf(d2))

}

func test(t interface{}) {
	switch reflect.TypeOf(t).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(t)

		for i := 0; i < s.Len(); i++ {
			fmt.Print(s.Index(i), " ")
		}
		fmt.Println()
	}
}
