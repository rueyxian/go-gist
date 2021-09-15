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

	if isLamer((*dog)(nil)) {
		fmt.Println("dog implements lame")
	}

	// ==============================

	if isLamer((*cat)(nil)) {
		fmt.Println("cat implements lame")
	}

	// ==============================

	if isLamer((*gopher)(nil)) {
		fmt.Println("gopher implements lame")
	}

}

// ================================================================================

func isLamer(v interface{}) bool {
	_, ok := v.(lamer)
	return ok
}
