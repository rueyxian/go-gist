package main

import (
	"fmt"
	"log"
)

type customError struct{}

func (c *customError) Error() string {
	return "custom error"
}

func operation() *customError {
	return nil
}

func main() {

	// reason...
	// https://golang.org/doc/faq#nil_error

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
