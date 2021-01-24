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

	// if we declare err as concreate type, the problem will be solved
	// but it's not the best practice
	var err *customError

	fmt.Printf("T: %[1]T      V: %[1]v\n", err)

	err = operation()

	fmt.Printf("T: %[1]T      V: %[1]v\n", err)

	if err != nil {
		log.Println("error!!")
		return
	}

	log.Println("have a nice day")

}
