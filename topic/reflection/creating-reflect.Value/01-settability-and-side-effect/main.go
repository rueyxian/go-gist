package main

import (
	"fmt"
	"reflect"
)

func main() {

	{
		fmt.Println("rv := ValueOf(T{})")

		v := "unchanged"
		rv := reflect.ValueOf(v)

		fmt.Println("\t", "CanSet():", rv.CanSet())
		fmt.Println("\t", "-")
		fmt.Println()
	}

	{
		fmt.Println("rv := ValueOf(&T{}).Elem()")

		v := "unchanged"
		rv := reflect.ValueOf(&v).Elem()
		rv.SetString("changed")

		fmt.Println("\t", "CanSet():", rv.CanSet())
		fmt.Println("\t", v)
		fmt.Println()
	}

	{
		fmt.Println("rv := ValueOf((*T)(nil)).Elem()")

		rv := reflect.ValueOf((*string)(nil)).Elem()

		fmt.Println("\t", "CanSet():", rv.CanSet())
		fmt.Println("\t", "-")
		fmt.Println()
	}

	// ==============================
	fmt.Println("==============================\n")
	// ==============================

	{
		fmt.Println("rt := TypeOf(T{})")
		fmt.Println("rv := New(rt).Elem()")

		v := "unchanged"
		rt := reflect.TypeOf(v)
		rv := reflect.New(rt).Elem()
		rv.SetString("changed")

		fmt.Println("\t", "CanSet():", rv.CanSet())
		fmt.Println("\t", v)
		fmt.Println()
	}

	{
		fmt.Println("rt := TypeOf(&T{}).Elem()")
		fmt.Println("rv := New(rt).Elem()")

		v := "unchanged"
		rt := reflect.TypeOf(&v).Elem()
		rv := reflect.New(rt).Elem()
		rv.SetString("changed")

		fmt.Println("\t", "CanSet():", rv.CanSet())
		fmt.Println("\t", v)
		fmt.Println()
	}

	{
		fmt.Println("rt := TypeOf((*T)(nil)).Elem()")
		fmt.Println("rv := New(rt).Elem()")

		var v *string
		rt := reflect.TypeOf(v).Elem()
		rv := reflect.New(rt).Elem()
		rv.SetString("changed")

		fmt.Println("\t", "CanSet():", rv.CanSet())
		fmt.Println("\t", v)
		fmt.Println()
	}

	// ==============================
	fmt.Println("==============================\n")
	// ==============================

	// {
	//   fmt.Println("rt := TypeOf(T{})")
	//   fmt.Println("rv := NewAt(rt).Elem()")
	//   v := T{}
	//   rt := reflect.TypeOf(T{})
	//   rv := reflect.NewAt(rv., unsafe.Pointer(rv.)).Elem()

	//   fmt.Println(rv.CanSet())
	//   fmt.Println()
	// }

	// {
	//   fmt.Println("rt := TypeOf(&T{}).Elem()")
	//   fmt.Println("rv := New(rt).Elem()")
	//   v := "original"
	//   rt := reflect.TypeOf(&v).Elem()
	//   rv := reflect.NewAt(rt, rv.).Elem()
	//   rv.SetString("changed")
	//   fmt.Println(v)
	//   // fmt.Println(rv.CanSet())
	//   fmt.Println()
	// }

	// {
	//   fmt.Println("rt := TypeOf((*T)(nil)).Elem()")
	//   fmt.Println("rv := New(rt).Elem()")
	//   rt := reflect.TypeOf((*T)(nil)).Elem()
	//   rv := reflect.New(rt).Elem()

	//   fmt.Println(rv.CanSet())
	//   fmt.Println()
	// }

}
