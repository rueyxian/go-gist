package main

import (
	"fmt"
	"strings"
)

func main() {

	{
		s := strings.SplitN("a,b,c,d,e,f,g", ",", 4)
		fmt.Printf("%q\n", s)
	}

	{
		fmt.Printf("%q\n", strings.SplitN("a,b,c", ",", 2))
	}

	{
		z := strings.SplitN("a,b,c", ",", 0)
		fmt.Printf("%q (nil = %v)\n", z, z == nil)
	}
}
