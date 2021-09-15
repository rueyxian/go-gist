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
	// Or, if we want it to be interface type ...
	var err error
	fmt.Printf("T=%[1]T | V=%[1]v\n", err)
	_, err = operation()
	fmt.Printf("T=%[1]T | V=%[1]v\n", err)

	// ... apply type convertion in the comparison statement:
	// checking if err is not equal to a pointer to nil custom error
	// However, this is by no means a good practice
	if err != (*customError)(nil) {
		fmt.Println(err)
		return
	}

	fmt.Println("peace out")
}
