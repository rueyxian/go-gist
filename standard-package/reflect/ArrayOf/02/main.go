package main

import (
	"fmt"
	"reflect"
)

func main() {
	ti := reflect.TypeOf(4896)

	ta := reflect.ArrayOf(4, ti)

	va := reflect.New(ta)

	ia := va.Interface()
	iaa := va.Interface().(*[4]int)

	fmt.Printf("%T  %v\n", ti, ti)
	fmt.Printf("%T  %v\n", ta, ta)
	fmt.Printf("%T  %v\n", va, va)

	fmt.Printf("%T  %v\n", ia, ia)
	fmt.Printf("%T  %v\n", iaa, iaa)

}

// func main() {
//     t := reflect.TypeOf(5)

//     // use of ArrayOf method
//     arr := reflect.ArrayOf(4, t)
//     inst := reflect.New(arr).Interface().(*[4]int)

//     for i := 1; i <= 4; i++ {
//         inst[i-1] = i*i
//     }

//     fmt.Println(inst)
// }
