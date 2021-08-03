package main

import (
	"fmt"
	"log"
)

func main() {
	err := operation()
	if err != nil {
		// but how come end up here?
		fmt.Println("how come end up here?")
		log.Fatal(err)
	}
	fmt.Println("end peacefully")
}

func operation() error {
	var p *myError

	// make sure it will not enter this block
	// so *myError == nil (always)
	if false {
		p = &myError{"error"}
	}
	return p
}

type myError struct {
	errMsg string
}

func (e myError) Error() string {
	return e.errMsg
}
