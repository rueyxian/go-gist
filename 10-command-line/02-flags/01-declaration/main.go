package main

import (
	"flag"
	"fmt"
)

func main() {

	// declare as pointer
	// - flag.String
	// - flag.Int
	// - flag.Bool
	greeting := flag.String("greeting", "Welcome", "startup message")
	counter := flag.Int("counter", 9, "counterrrr")

	// bind with a variable
	// - flag.StringVar
	// - flag.IntVar
	// - flag.BoolVar
	var favLang string
	flag.StringVar(&favLang, "favLang", "Go", "favourite language")

	fmt.Println("greeting: ", *greeting)
	fmt.Println("counter: ", *counter)
	fmt.Println("favLang: ", favLang)

}
