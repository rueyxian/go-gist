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
	var err error // err is declared as interface value
	fmt.Printf("T=%[1]T | V=%[1]v\n", err)
	_, err = operation()
	fmt.Printf("T=%[1]T | V=%[1]v\n", err)

	if err != nil {
		// println("ERROR!!!! ARRG!")
		fmt.Println(err)
		return
	}

	fmt.Println("peace out")
}

// It returns error. Why?
// Because err is decalred as interface value,
// eventhough operation() returns as nil
// err != nil as schematically, (T = error, V = nil).
// An interface value is nil only if both T & V are nil.

// golang.org faq explains very clearly:
// Why is my nil error value not equal to nil?
// https://golang.org/doc/faq#nil_error
