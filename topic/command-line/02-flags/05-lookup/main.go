package main

import (
	"flag"
	"fmt"
)

var testFlag = flag.String("test", "default value", "help message.")

func print(f *flag.Flag) {
	if f != nil {
		fmt.Println(f.Value)
	} else {
		fmt.Println(nil)
	}
}

func main() {

	fmt.Print("test:")
	print(flag.Lookup("test"))
	fmt.Print("test1:")
	print(flag.Lookup("test1"))

	flag.Parse()
	fmt.Print("test:")
	print(flag.Lookup("test"))
	fmt.Print("test1:")
	print(flag.Lookup("test1"))
}
