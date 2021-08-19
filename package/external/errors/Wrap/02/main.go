package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func main() {

	if err := run(); err != nil {
		fmt.Println(err)
	}

}

func run() error {
	if err := cause(); err != nil {
		return errors.Wrap(err, "oh no!")
	}
	return nil
}

func cause() error {
	return errors.New("whoops")
}
