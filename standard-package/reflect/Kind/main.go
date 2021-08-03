package main

import (
	"fmt"
	"reflect"
)

func main() {

	s := []interface{}{"hi", 42, func() {}}

	withoutReflect(s)
	fmt.Println()
	withReflect(s)

}

// ================================================================================

func withoutReflect(s []interface{}) {
	for _, v := range s {
		switch v := v.(type) {
		case string:
			fmt.Println(v)
		case int, int8, int16, int32, int64:
			fmt.Println(v)
		default:
			fmt.Printf("unhandled kind %s\n", v)
		}
	}
}

func withReflect(s []interface{}) {
	for _, v := range s {
		switch v := reflect.ValueOf(v); v.Kind() {
		case reflect.String:
			fmt.Println(v.String())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			fmt.Println(v.Int())
		default:
			fmt.Printf("unhandled kind %s\n", v.Kind())
		}
	}
}

// ================================================================================
