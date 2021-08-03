package main

import (
	"fmt"
)

type customError struct{}

func (c *customError) Error() string {
	if c == nil {
		return "custom nil!"
	}
	return "custom error!"
}

func operation() ([]byte, *customError) {
	// rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	// if rnd.Intn(2) == 0 {
	//   return nil, &customError{}
	// }
	// return []byte("some unimportant data"), nil
	return nil, nil
}

func main() {
	var err error
	fmt.Printf("err: [ T=%[1]T, V=%[1]v ]\n", err)
	_, err = operation()
	fmt.Printf("err: [ T=%[1]T, V=%[1]v ]\n", err)

	if v, ok := err.(*customError); ok && v != nil {
		println("ERROR!!!! ARRG!")
		return
	}

	println("peace out")
}
