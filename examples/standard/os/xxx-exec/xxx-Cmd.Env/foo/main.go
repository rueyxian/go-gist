package main

import (
	"flag"
	"fmt"
)

func main() {

	name := flag.String("name", "luna", "a string")
	age := flag.Int("age", 21, "an int")
	flag.Parse()

	fmt.Printf("name: %s | age: %d\n", *name, *age)

}
