package main

import (
	"fmt"

	"github.com/pkg/errors"
)

type AppError struct {
	State int
}

func (c *AppError) Error() string {
	return fmt.Sprintf("App Error, State: %d", c.State)
}

func main() {

	if err := callA(123); err != nil {

		// switch v := err.(type) {
		switch v := errors.Cause(err).(type) {
		case *AppError:

			// We got our custom error type.
			fmt.Println("\nRoot\n********************************")
			fmt.Printf("[root] Custom App Error: %v\n", v.State)

		default:

			// We did not get any specific error type.
			fmt.Println("Default Error")
		}

		fmt.Println("\nStack Trace\n********************************")
		fmt.Printf("%+v\n", err)
		fmt.Println()
		fmt.Println("\nNo Trace\n********************************")
		fmt.Printf("%v\n", err)
	}

}

func callA(i int) error {
	if err := callB(99); err != nil {
		return errors.Wrapf(err, "callA(%v) -> callB(%v)", i, i)
	}
	return nil
}

func callB(i int) error {
	if err := callC(); err != nil {
		return errors.Wrapf(err, "callB(%v) -> callC()", i)
	}
	return nil
}

func callC() error {
	return &AppError{666}
}
