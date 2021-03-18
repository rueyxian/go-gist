package main

import (
	"fmt"
	"reflect"
)

type A [16]int16

// ========================================

// type T []interface{}
type T interface {
	n()
}

// func (T) n() {}

// ========================================
type U []interface {
	m()
}

func (U) m() {}

// ========================================

func main() {

	var t T
	var u U

	tt := reflect.TypeOf(t)
	fmt.Println(tt)
	// fmt.Println(tt.Kind())
	fmt.Println()

	tu := reflect.TypeOf(u)
	fmt.Println(tu)
	// fmt.Println(tu.Kind())

	// ========================================
	fmt.Println()
	fmt.Println("============================================================")
	fmt.Println()
	// ========================================

	var v interface{}
	p := new(interface{})

	tv := reflect.TypeOf(v)
	tp := reflect.TypeOf(p)

	fmt.Println(tv)
	// fmt.Println(tv.Kind())
	fmt.Println()

	fmt.Println(tp)
	fmt.Println(tp.Kind())

	// ========================================
	fmt.Println()
	fmt.Println("============================================================")
	fmt.Println()
	// ========================================

	var i *int
	ti := reflect.TypeOf(i)
	fmt.Println(ti)
	fmt.Println(ti.Kind())
	fmt.Println(ti.Elem())

}
