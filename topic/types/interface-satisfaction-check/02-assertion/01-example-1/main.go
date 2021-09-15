package main

import "fmt"

// ================================================================================

type lamer interface {
	lame() string
}

// ================================================================================

type dog struct{}

func (_ dog) lame() string {
	return "I'm so lame"
}

// ================================================================================

type cat string

func (_ cat) lame() string {
	return "I'm freaking awesome"
}

// ================================================================================

type gopher int

func (_ gopher) awesome() string {
	return "I'm awesome too"
}

// ================================================================================

func main() {

	var v1 interface{} = (*dog)(nil)

	if _, ok := v1.(lamer); ok {
		fmt.Println("dog implements lame")
	}

	// ==============================

	var v2 interface{} = (*cat)(nil)

	if _, ok := v2.(lamer); ok {
		fmt.Println("cat implements lame")
	}

	// ==============================

	var v3 interface{} = (*gopher)(nil)

	if _, ok := v3.(lamer); ok {
		fmt.Println("gopher implements lame")
	}

}
