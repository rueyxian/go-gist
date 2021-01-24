package main

type customError struct{}

func (c *customError) Error() string {
	if c == nil {
		return "custom nil!"
	}
	return "custom error!"
}

func operation() ([]byte, *customError) {
	return nil, nil
	//return nil, &customError //how to return a custom error
}

func main() {
	var err error
	_, err = operation()

	if err != (*customError)(nil) { //here we are checking if err is not equal to a pointer to nil custom error
		println("ERROR!!!! ARRG!")
		return
	}

	println("peace out")
}
