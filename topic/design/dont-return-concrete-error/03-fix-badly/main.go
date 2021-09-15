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

	// To combat this, we need to predeclare it as *customError,
	// a concrete type instead of interface type
	//  _, err := operation()   <-- works the same
	var err *customError
	fmt.Printf("T=%[1]T | V=%[1]v\n", err)
	_, err = operation()
	fmt.Printf("T=%[1]T | V=%[1]v\n", err)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("peace out")
}
