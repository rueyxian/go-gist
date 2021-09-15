package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {

	// To combat this, we need to add an extra layer of function.
	// The rule of thumb is: only call os.Exit() – or function/method
	// that contains os.Exit() – at the highest level call stack,
	// which is main func

	if err := run(); err != nil {
		log.Fatal(err)
	}

}

func run() error {
	defer important()
	if err := execute(); err != nil {
		return err
	}
	return nil
}

func execute() error {
	return errors.New("ouch!")
}

func important() {
	fmt.Println("something important")
}
