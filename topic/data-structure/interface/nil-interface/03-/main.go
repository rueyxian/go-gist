package main

import "fmt"

type customError struct{}

func (c *customError) Error() string {
	return "custom error!"
}

func operation() ([]byte, *customError) {
	return nil, nil
}

func main() {
	var err error

	// fmt.Println(err) //here err is a nil pointer to error interface
	fmt.Printf("[%[1]T, %[1]v]\n", err) //here err is a nil pointer to error interface

	_, err = operation()

	fmt.Printf("[%[1]T, %[1]v]\n", err) //here err is pointer to nil customError interface, err is not nil itself

	if err != nil {
		println("ERROR!!!! ARRG!")
		return
	}

	println("peace out")
}
