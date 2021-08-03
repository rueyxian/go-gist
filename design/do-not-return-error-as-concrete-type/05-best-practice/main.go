package main

import "fmt"

type customError struct{}

func (c *customError) Error() string {
	return "custom error!"
}

// Most of the time, we shouldn't return interface type
// However, when dealing with error handling, it's another way round
// Return concrete type is a bad practice
// We should always return error (interface type)
// to avoid unwanted mistake as previous example shown
func operation() ([]byte, error) {
	return nil, nil
}

func main() {
	var err error // it doesn't mater if you pre declare err or not
	fmt.Printf("T=%[1]T | V=%[1]v\n", err)
	_, err = operation()
	fmt.Printf("T=%[1]T | V=%[1]v\n", err)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("peace out")

}
