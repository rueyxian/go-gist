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
	_, err := operation() // err is declared as concrete value
	fmt.Printf("T=%[1]T | V=%[1]v\n", err)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("peace out")
}

// peace out will be printed, as expected
