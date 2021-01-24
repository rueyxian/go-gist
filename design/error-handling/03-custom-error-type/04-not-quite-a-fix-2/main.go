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

	_, k := err.(*customError)
	fmt.Println("lol   ", k)

	// same thing here, it's not a reasonable solution
	// it simply assert err as *customError, and check whether it is nil
	// we assumed err can be asserted as *customError, what if it can't
	if v, ok := err.(*customError); ok && v != nil {
		log.Println("error!!")
		return
	}

	log.Println("have a nice day")

}
