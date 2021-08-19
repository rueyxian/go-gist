package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {

	// if we look at log.Fatal() definition, it calls os.Exit(1)
	// at the end, which means any deferred code will be skipped

	// func Fatal(v ...interface{}) {
	//   std.Output(2, fmt.Sprint(v...))
	//   os.Exit(1)
	// }

	defer important()

	if err := execute(); err != nil {
		log.Fatal(err)
	}

}

func execute() error {
	return errors.New("ouch!")
}

func important() {
	fmt.Println("something important")
}
