package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func main() {
	cause := errors.New("whoops")
	err := errors.Wrap(cause, "oh noes")

	fmt.Println(err)

}
