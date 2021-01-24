package main

import (
	"fmt"
	"log"
)

type customError struct{}

func (c *customError) Error() string {
	return "custom error"
}

// we should avoid returning concreate error type for error handling
// return error (interface type is the best practice)
func operation() error {
	return nil
}

func main() {

	var err error

	fmt.Printf("T: %[1]T      V: %[1]v\n", err)

	err = operation()

	fmt.Printf("T: %[1]T      V: %[1]v\n", err)

	if err != nil {
		log.Println("error!!")
		return
	}

	log.Println("have a nice day")

}
