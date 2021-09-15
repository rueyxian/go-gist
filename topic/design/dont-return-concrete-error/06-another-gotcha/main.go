package main

import "fmt"

type customError struct{}

func (c *customError) Error() string {
	return "custom error!"
}

func operation() ([]byte, error) {
	// When we predeclare it as concrete type,
	// then return it as interface type,
	// it will end up being interface type with T=*customError
	var err *customError = nil
	return nil, err
}

func main() {
	_, err := operation()
	fmt.Printf("T=%[1]T | V=%[1]v\n", err)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("peace out")

}
