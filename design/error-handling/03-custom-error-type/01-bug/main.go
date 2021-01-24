package main

import "log"

type customError struct{}

func (c *customError) Error() string {
	return "custom error"
}

func operation() error {
	var err *customError
	//err == nil
	return err
}

func main() {

	err := operation()

	if err != nil {
		log.Println("error!!")
		return
	}

	log.Println("have a nice day")

}
