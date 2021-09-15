package main

import (
	"fmt"
	"reflect"
)

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

	i := reflect.TypeOf((*lamer)(nil)).Elem()
	v1 := reflect.TypeOf((*dog)(nil)).Elem()
	v2 := reflect.TypeOf((*cat)(nil)).Elem()
	v3 := reflect.TypeOf((*gopher)(nil)).Elem()

	// ==============================

	if v1.Implements(i) {
		fmt.Println("dog implements lame")
	}

	// ==============================

	if v2.Implements(i) {
		fmt.Println("cat implements lame")
	}

	// ==============================

	if v3.Implements(i) {
		fmt.Println("gopher implements lame")
	}

}
