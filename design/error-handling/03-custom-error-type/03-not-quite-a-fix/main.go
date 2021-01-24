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

	var err error

	fmt.Printf("T: %[1]T      V: %[1]v\n", err)

	err = operation()

	fmt.Printf("T: %[1]T      V: %[1]v\n", err)

	// this is not a solution though
	// it just simply check whether err is customError nil, not error nil
	if err != (*customError)(nil) {
		log.Println("error!!")
		return
	}

	log.Println("have a nice day")

}
