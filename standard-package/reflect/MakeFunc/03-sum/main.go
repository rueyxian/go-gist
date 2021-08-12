package main

import (
	"fmt"
	"reflect"
)

func main() {

	sum := func(args []reflect.Value) []reflect.Value {
		a, b := args[0], args[1]
		if a.Kind() != b.Kind() {
			return nil
		}

		switch a.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return []reflect.Value{reflect.ValueOf(a.Int() + b.Int())}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return []reflect.Value{reflect.ValueOf(a.Uint() + b.Uint())}
		case reflect.Float32, reflect.Float64:
			return []reflect.Value{reflect.ValueOf(a.Float() + b.Float())}
		case reflect.String:
			return []reflect.Value{reflect.ValueOf(a.String() + b.String())}
		default:
			return []reflect.Value{}
		}

	}

	makeSum := func(fptr interface{}) {
		rv := reflect.ValueOf(fptr).Elem()
		rv.Set(reflect.MakeFunc(rv.Type(), sum))
	}

	var sumInt func(int, int) int64
	var sumFloat func(float64, float64) float64
	var sumString func(string, string) string
	// var sumError func(string, int) string

	makeSum(&sumInt)
	makeSum(&sumFloat)
	makeSum(&sumString)
	// makeSum(&sumError)

	fmt.Println(sumInt(3, 5))
	fmt.Println(sumFloat(1.61, 3.14))
	fmt.Println(sumString("lfg", "2000"))
	// fmt.Println(sumError("bfg", 9000))

}
