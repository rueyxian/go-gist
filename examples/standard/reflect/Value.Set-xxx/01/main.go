package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// func main() {

//   var v = struct {
//     foo string
//     bar int
//   }{"dumb", 7}
//   // v0 := "bfg"
//   // v1 := 9000

//   rv := reflect.ValueOf(&v).Elem()
//   rf0 := rv.Field(0)
//   rf1 := rv.Field(1)

//   rv0 := reflect.ValueOf(stringPtr("bfg")).Elem()
//   rv1 := reflect.ValueOf(intPtr(9000)).Elem()

//   rf0.Set(rv0)
//   rf1.Set(rv1)

//   // fmt.Printf("%[1]T | %[1]v\n", rv)
//   // fmt.Printf("%[1]T | %[1]v\n", rv0)
//   // fmt.Printf("%[1]T | %[1]v\n", rv1)
//   fmt.Printf("%[1]T | %[1]v\n", rv)
//   // rf0 := rs.Field(0)

// }

// ================================================================================

func main() {
	var s = struct{ foo int }{654}
	var i int = 99

	rs := reflect.ValueOf(&s).Elem()
	rf := rs.Field(0)
	ri := reflect.ValueOf(&i).Elem()

	rf = reflect.NewAt(rf.Type(), unsafe.Pointer(rf.UnsafeAddr())).Elem()
	// ri.Set(rf)
	rf.Set(ri)
	fmt.Println(rf)
	fmt.Println(ri)
	fmt.Println(rs)

}

// ================================================================================

// type stringLitr string

func stringPtr(s string) *string {
	return &s
}

func intPtr(i int) *int {
	return &i
}
