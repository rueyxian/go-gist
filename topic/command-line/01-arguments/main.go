package main

import (
	"fmt"
	"os"
)

func main() {

	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	fmt.Printf("%T \t %+v \n", argsWithProg, argsWithProg)
	fmt.Printf("%T \t %+v \n", argsWithoutProg, argsWithoutProg)

}
